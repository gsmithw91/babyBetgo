#!/usr/bin/env python3

import os
import sys
import subprocess

MIGRATIONS_DIR = "database/migrations"
DEFAULT_DB_URL = "postgres://babybetgo_user:password@localhost:5432/babybetgo_dev?sslmode=disable"

DATABASE_URL = os.getenv("DATABASE_URL", DEFAULT_DB_URL)


def run_migrate_command(*xargs):
    cmd = [
        "migrate",
        "-path", MIGRATIONS_DIR,
        "-database", DATABASE_URL,
        *xargs
    ]

    print(f"🔧 Running: {' '.join(cmd)}")
    subprocess.run(cmd, check=True)


def main():
    if len(sys.argv) < 2:
        print("Usage: migrate.py [up|down|drop|version|create <name>]")
        return

    command = sys.argv[1]

    try:
        match command:
            case "up":
                run_migrate_command("up")
            case "down":
                run_migrate_command("down", "1")
            case "drop":
                run_migrate_command("drop", "-f")
            case "version":
                run_migrate_command("version")
            case "create":
                if len(sys.argv) < 3:
                    print("❌ Missing name for migration: migrate.py create <name>")
                    return
                name = sys.argv[2]
                subprocess.run([
                    "migrate", "create", "-ext", "sql", "-dir", MIGRATIONS_DIR, "-seq", name
                ], check=True)
            case _:
                print(f"❌ Unknown command: {command}")
    except subprocess.CalledProcessError as e:
        print(f"❌ Migration command failed: {e}")


if __name__ == "__main__":
    main()

