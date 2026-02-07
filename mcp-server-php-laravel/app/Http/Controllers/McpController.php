<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use App\Services\ToolRegistry;

class McpController extends Controller
{
    protected $toolRegistry;

    public function __construct(ToolRegistry $toolRegistry)
    {
        $this->toolRegistry = $toolRegistry;
    }

    /**
     * Root endpoint
     * GET /
     */
    public function index(): JsonResponse
    {
        return response()->json([
            'message' => '🧠 MCP Server PHP Laravel is running'
        ]);
    }

    /**
     * Health check endpoint
     * GET /health
     */
    public function health(): JsonResponse
    {
        return response()->json([
            'status' => 'ok'
        ]);
    }

    /**
     * List all available tools
     * GET /mcp/tools
     */
    public function tools(): JsonResponse
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

        return response()->json($toolsResponse);
    }

    /**
     * Invoke a tool
     * POST /mcp/invoke
     */
    public function invoke(Request $request): JsonResponse
    {
        $toolName = $request->input('name');
        $arguments = $request->input('arguments', []);

        if (!$toolName) {
            return response()->json(['error' => 'Tool name is required'], 400);
        }

        $tool = $this->toolRegistry->findTool($toolName);

        if (!$tool) {
            return response()->json(['error' => 'Tool not found'], 404);
        }

        if (!isset($tool['handler']) || !is_callable($tool['handler'])) {
            return response()->json(['error' => 'Tool handler not found'], 400);
        }

        try {
            $result = call_user_func($tool['handler'], $arguments);
            return response()->json(['result' => $result]);
        } catch (\Exception $e) {
            return response()->json(['error' => $e->getMessage()], 500);
        }
    }
}
