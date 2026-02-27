<?php

namespace CodeSmells\Couplers;

class Database
{
    public function query(string $sql): array
    {
        echo "Executing query: $sql\n";
        return [];
    }

    public function insert(string $table, array $data): void
    {
        echo "Inserting into $table\n";
    }

    public function update(string $table, array $data): void
    {
        echo "Updating $table\n";
    }
}

class DataManager
{
    private Database $database;

    public function __construct(Database $database)
    {
        $this->database = $database;
    }

    public function query(string $sql): array
    {
        return $this->database->query($sql);
    }

    public function insert(string $table, array $data): void
    {
        $this->database->insert($table, $data);
    }

    public function update(string $table, array $data): void
    {
        $this->database->update($table, $data);
    }
}
