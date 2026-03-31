#!/usr/bin/env python3
import logging
import os
from dataclasses import dataclass

from dotenv import load_dotenv

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class AppConfig:
    """env: dev | prod
    """

    env: str  # "dev" | "prod"

    @classmethod
    def load(cls) -> "AppConfig":
        if os.getenv("ENV") != "prod":
            if load_dotenv():
                logger.info(".env loaded")
            else:
                logger.warning("no .env file found")

        required = ["ENV"]
        missing = [k for k in required if not os.getenv(k)]
        if missing:
            raise RuntimeError(f"missing required environment variables: {', '.join(missing)}")

        return cls(env=os.getenv("ENV", "dev"))

    @property
    def is_dev(self) -> bool:
        return self.env == "dev"

    @property
    def is_prod(self) -> bool:
        return self.env == "prod"
