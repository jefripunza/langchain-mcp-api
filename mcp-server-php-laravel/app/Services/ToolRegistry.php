<?php

namespace App\Services;

use App\Services\Tools\StringTools;
use App\Services\Tools\MathTools;
use App\Services\Tools\NetworkTools;
use App\Services\Tools\FileTools;

class ToolRegistry
{
    private $tools = [];

    public function __construct()
    {
        $this->registerTools();
    }

    private function registerTools()
    {
        // Register String Tools
        $stringTools = new StringTools();
        foreach ($stringTools->getTools() as $tool) {
            $this->tools[] = $tool;
        }

        // Register Math Tools
        $mathTools = new MathTools();
        foreach ($mathTools->getTools() as $tool) {
            $this->tools[] = $tool;
        }

        // Register Network Tools
        $networkTools = new NetworkTools();
        foreach ($networkTools->getTools() as $tool) {
            $this->tools[] = $tool;
        }

        // Register File Tools
        $fileTools = new FileTools();
        foreach ($fileTools->getTools() as $tool) {
            $this->tools[] = $tool;
        }
    }

    public function getTools()
    {
        return $this->tools;
    }

    public function findTool($name)
    {
        foreach ($this->tools as $tool) {
            if ($tool['name'] === $name) {
                return $tool;
            }
        }
        return null;
    }
}
