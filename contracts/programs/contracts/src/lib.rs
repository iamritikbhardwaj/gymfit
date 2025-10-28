use anchor_lang::prelude::*;

declare_id!("FqRSiop8ovVwTb9zT7eFU9r43835FizLWyaoR8CHg88R");

#[program]
pub mod contracts {
    use super::*;

    pub fn initialize(ctx: Context<Initialize>) -> Result<()> {
        msg!("Greetings from: {:?}", ctx.program_id);
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize {}
