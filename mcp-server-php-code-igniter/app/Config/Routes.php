<?php

use CodeIgniter\Router\RouteCollection;

/**
 * @var RouteCollection $routes
 */
// MCP Server Routes
$routes->get('/', 'McpController::index');
$routes->get('/health', 'McpController::health');
$routes->get('/mcp/tools', 'McpController::tools');
$routes->post('/mcp/invoke', 'McpController::invoke');
