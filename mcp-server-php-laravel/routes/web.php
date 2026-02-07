<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\McpController;

// MCP Server Routes
Route::get('/', [McpController::class, 'index']);
Route::get('/health', [McpController::class, 'health']);
Route::get('/mcp/tools', [McpController::class, 'tools']);
Route::post('/mcp/invoke', [McpController::class, 'invoke']);
