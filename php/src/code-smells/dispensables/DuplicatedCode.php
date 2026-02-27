<?php

namespace CodeSmells\Dispensables;

class ReportA
{
    public function generateReport(array $data): string
    {
        $report = "Report A\n";
        $report .= "==============\n";
        foreach ($data as $item) {
            $report .= "- {$item}\n";
        }
        $report .= "==============\n";
        return $report;
    }
}

class ReportB
{
    public function generateReport(array $data): string
    {
        $report = "Report B\n";
        $report .= "==============\n";
        foreach ($data as $item) {
            $report .= "- {$item}\n";
        }
        $report .= "==============\n";
        return $report;
    }
}
