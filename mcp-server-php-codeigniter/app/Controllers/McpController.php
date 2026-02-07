<?php

namespace App\Controllers;

use CodeIgniter\RESTful\ResourceController;
use App\Libraries\ToolRegistry;

class McpController extends ResourceController
{
    protected $format = 'json';
    protected $toolRegistry;

    public function __construct()
    {
        $this->toolRegistry = new ToolRegistry();
    }

    /**
     * Root endpoint
     * GET /
     */
    public function index()
    {
        return $this->respond([
            'message' => '🧠 MCP Server PHP CodeIgniter is running'
        ]);
    }

    /**
     * Health check endpoint
     * GET /health
     */
    public function health()
    {
        return $this->respond([
            'status' => 'ok'
        ]);
    }

    /**
     * List all available tools
     * GET /mcp/tools
     */
    public function tools()
    {
        $tools = $this->toolRegistry->getTools();
        $toolsResponse = [];

        foreach ($tools as $tool) {
            $toolsResponse[] = [
                'name' => $tool['name'],
                'description' => $tool['description'],
                'parameters' => $tool['parameters']
            ];
        }

        return $this->respond($toolsResponse);
    }

    /**
     * Invoke a tool
     * POST /mcp/invoke
     */
    public function invoke()
    {
        $json = $this->request->getJSON(true);

        if (!isset($json['name'])) {
            return $this->fail('Tool name is required', 400);
        }

        $toolName = $json['name'];
        $arguments = $json['arguments'] ?? [];

        $tool = $this->toolRegistry->findTool($toolName);

        if (!$tool) {
            return $this->fail('Tool not found', 404);
        }

        if (!isset($tool['handler']) || !is_callable($tool['handler'])) {
            return $this->fail('Tool handler not found', 400);
        }

        try {
            $result = call_user_func($tool['handler'], $arguments);
            return $this->respond(['result' => $result]);
        } catch (\Exception $e) {
            return $this->fail($e->getMessage(), 500);
        }
    }
}
