use anyhow::Context as _;
use async_trait::async_trait;
use futures::{Stream, StreamExt};
use sqlx::Either;
use std::{ops::Range, pin::Pin};

use penumbra_tct::{
    storage::{AsyncRead, AsyncWrite, StoredPosition},
    structure::Hash,
    Commitment, Forgotten, Position,
};

pub struct TreeStore<'a, 'c: 'a>(pub &'a mut sqlx::Transaction<'c, sqlx::Sqlite>);

#[async_trait]
impl AsyncRead for TreeStore<'_, '_> {
    type Error = anyhow::Error;

    async fn position(&mut self) -> Result<StoredPosition, Self::Error> {
        Ok(sqlx::query!("SELECT position FROM sct_position LIMIT 1")
            .fetch_one(&mut *self.0)
            .await?
            .position
            .map(|p| Position::from(p as u64))
            .into())
    }

    async fn forgotten(&mut self) -> Result<Forgotten, Self::Error> {
        Ok((sqlx::query!("SELECT forgotten FROM sct_forgotten LIMIT 1")
            .fetch_one(&mut *self.0)
            .await?
            .forgotten as u64)
            .into())
    }

    async fn hash(&mut self, position: Position, height: u8) -> Result<Option<Hash>, Self::Error> {
        let position = u64::from(position) as i64;

        let bytes = sqlx::query!(
            "SELECT hash FROM sct_hashes WHERE position = ? AND height = ? LIMIT 1",
            position,
            height
        )
        .fetch_optional(&mut *self.0)
        .await?
        .map(|r| r.hash);

        Ok(bytes
            .map(|bytes| {
                <[u8; 32]>::try_from(bytes)
                    .map_err(|_| anyhow::anyhow!("hash was of incorrect length"))
                    .and_then(|array| Hash::from_bytes(array).map_err(Into::into))
            })
            .transpose()?)
    }

    fn hashes(
        &mut self,
    ) -> Pin<Box<dyn Stream<Item = Result<(Position, u8, Hash), Self::Error>> + Send + '_>> {
        Box::pin(
            sqlx::query!("SELECT position, height, hash FROM sct_hashes")
                .fetch_many(&mut *self.0)
                .map(|row| {
                    let row = row?;
                    if let Either::Right(row) = row {
                        Ok::<_, Self::Error>(Some((
                            Position::from(row.position as u64),
                            row.height as u8,
                            Hash::from_bytes(
                                row.hash
                                    .try_into()
                                    .map_err(|_| anyhow::anyhow!("hash was of incorrect length"))?,
                            )?,
                        )))
                        .context("could not decode hash from local database")
                    } else {
                        Ok(None)
                    }
                })
                .filter_map(|item| async move { item.transpose() }),
        )
    }

    async fn commitment(&mut self, position: Position) -> Result<Option<Commitment>, Self::Error> {
        let position = u64::from(position) as i64;

        let bytes: Option<Vec<u8>> = sqlx::query!(
            "SELECT commitment FROM sct_commitments WHERE position = ? LIMIT 1",
            position,
        )
        .fetch_optional(&mut *self.0)
        .await?
        .map(|r| r.commitment);

        Ok(bytes
            .map(|bytes| {
                <[u8; 32]>::try_from(bytes)
                    .map_err(|_| anyhow::anyhow!("commitment was of incorrect length"))
                    .and_then(|array| Commitment::try_from(array).map_err(Into::into))
            })
            .transpose()?)
    }

    fn commitments(
        &mut self,
    ) -> Pin<Box<dyn Stream<Item = Result<(Position, Commitment), Self::Error>> + Send + '_>> {
        Box::pin(
            sqlx::query!("SELECT position, commitment FROM sct_commitments")
                .fetch_many(&mut *self.0)
                .map(|row| {
                    let row = row?;
                    if let Either::Right(row) = row {
                        Ok::<_, Self::Error>(Some((
                            Position::from(row.position as u64),
                            Commitment::try_from(<[u8; 32]>::try_from(row.commitment).map_err(
                                |_| anyhow::anyhow!("commitment was of incorrect length"),
                            )?)?,
                        )))
                        .context("could not decode note commitment from local database")
                    } else {
                        Ok(None)
                    }
                })
                .filter_map(|item| async move { item.transpose() }),
        )
    }
}

#[async_trait]
impl AsyncWrite for TreeStore<'_, '_> {
    async fn set_position(&mut self, position: StoredPosition) -> Result<(), Self::Error> {
        let position = Option::from(position).map(|p: Position| u64::from(p) as i64);
        sqlx::query!("UPDATE sct_position SET position = ?", position)
            .execute(&mut *self.0)
            .await?;
        Ok(())
    }

    async fn set_forgotten(&mut self, forgotten: Forgotten) -> Result<(), Self::Error> {
        let forgotten = u64::from(forgotten) as i64;
        sqlx::query!("UPDATE sct_forgotten SET forgotten = ?", forgotten)
            .execute(&mut *self.0)
            .await?;
        Ok(())
    }

    async fn add_hash(
        &mut self,
        position: Position,
        height: u8,
        hash: Hash,
        _essential: bool,
    ) -> Result<(), Self::Error> {
        let position = u64::from(position) as i64;
        let hash = hash.to_bytes().to_vec();

        sqlx::query!(
            "INSERT INTO sct_hashes (position, height, hash) VALUES (?, ?, ?) ON CONFLICT DO NOTHING",
            position,
            height,
            hash
        )
        .execute(&mut *self.0).await?;
        Ok(())
    }

    async fn add_commitment(
        &mut self,
        position: Position,
        commitment: Commitment,
    ) -> Result<(), Self::Error> {
        let position = u64::from(position) as i64;
        let commitment = <[u8; 32]>::from(commitment).to_vec();
        sqlx::query!(
            "INSERT INTO sct_commitments (position, commitment) VALUES (?, ?) ON CONFLICT DO NOTHING",
            position,
            commitment
        )
        .execute(&mut *self.0)
        .await?;
        Ok(())
    }

    async fn delete_range(
        &mut self,
        below_height: u8,
        positions: Range<Position>,
    ) -> Result<(), Self::Error> {
        let start = u64::from(positions.start) as i64;
        let end = u64::from(positions.end) as i64;
        sqlx::query!(
            "DELETE FROM sct_hashes WHERE position >= ? AND position < ? AND height < ?",
            start,
            end,
            below_height
        )
        .execute(&mut *self.0)
        .await?;
        Ok(())
    }
}
