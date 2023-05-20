package main

//go:generate sqlboiler --config ./configs/sqlboiler.master.toml --wipe mysql
//go:generate sqlboiler --config ./configs/sqlboiler.shard.toml --wipe mysql
