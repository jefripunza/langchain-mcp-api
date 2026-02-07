<?php

namespace App\Services\Tools;

class NetworkTools
{
    public function getTools()
    {
        return [
            [
                'name' => 'network_validate_ip',
                'description' => 'Validasi IP address (IPv4/IPv6)',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'ip' => [
                            'type' => 'string',
                            'description' => 'IP address yang akan divalidasi'
                        ]
                    ],
                    'required' => ['ip']
                ],
                'handler' => [$this, 'validateIp']
            ],
            [
                'name' => 'network_dns_lookup',
                'description' => 'DNS lookup untuk mendapatkan IP dari hostname',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'hostname' => [
                            'type' => 'string',
                            'description' => 'Hostname yang akan di-lookup'
                        ]
                    ],
                    'required' => ['hostname']
                ],
                'handler' => [$this, 'dnsLookup']
            ],
            [
                'name' => 'network_parse_url',
                'description' => 'Parse URL menjadi komponen-komponennya',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'url' => [
                            'type' => 'string',
                            'description' => 'URL yang akan di-parse'
                        ]
                    ],
                    'required' => ['url']
                ],
                'handler' => [$this, 'parseUrl']
            ]
        ];
    }

    public function validateIp($args)
    {
        $ip = $args['ip'] ?? '';
        
        if (filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_IPV4)) {
            return [
                'valid' => true,
                'version' => 4,
                'type' => 'IPv4'
            ];
        } elseif (filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_IPV6)) {
            return [
                'valid' => true,
                'version' => 6,
                'type' => 'IPv6'
            ];
        } else {
            return [
                'valid' => false,
                'error' => 'Invalid IP address'
            ];
        }
    }

    public function dnsLookup($args)
    {
        $hostname = $args['hostname'] ?? '';
        
        $ip = gethostbyname($hostname);
        
        if ($ip === $hostname) {
            return ['error' => 'DNS lookup failed'];
        }
        
        return [
            'hostname' => $hostname,
            'ip' => $ip
        ];
    }

    public function parseUrl($args)
    {
        $url = $args['url'] ?? '';
        
        $parsed = parse_url($url);
        
        if ($parsed === false) {
            return ['error' => 'Invalid URL'];
        }
        
        return [
            'scheme' => $parsed['scheme'] ?? null,
            'host' => $parsed['host'] ?? null,
            'port' => $parsed['port'] ?? null,
            'path' => $parsed['path'] ?? null,
            'query' => $parsed['query'] ?? null,
            'fragment' => $parsed['fragment'] ?? null
        ];
    }
}
