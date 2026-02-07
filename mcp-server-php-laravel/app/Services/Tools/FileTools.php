<?php

namespace App\Services\Tools;

class FileTools
{
    public function getTools()
    {
        return [
            [
                'name' => 'file_get_extension',
                'description' => 'Dapatkan ekstensi file dari nama file',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'filename' => [
                            'type' => 'string',
                            'description' => 'Nama file'
                        ]
                    ],
                    'required' => ['filename']
                ],
                'handler' => [$this, 'getExtension']
            ],
            [
                'name' => 'file_get_basename',
                'description' => 'Dapatkan nama file tanpa path',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'path' => [
                            'type' => 'string',
                            'description' => 'Path lengkap file'
                        ]
                    ],
                    'required' => ['path']
                ],
                'handler' => [$this, 'getBasename']
            ],
            [
                'name' => 'file_format_bytes',
                'description' => 'Format ukuran bytes ke format yang lebih readable (KB, MB, GB)',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'bytes' => [
                            'type' => 'number',
                            'description' => 'Ukuran dalam bytes'
                        ]
                    ],
                    'required' => ['bytes']
                ],
                'handler' => [$this, 'formatBytes']
            ]
        ];
    }

    public function getExtension($args)
    {
        $filename = $args['filename'] ?? '';
        $extension = pathinfo($filename, PATHINFO_EXTENSION);
        return ['extension' => $extension];
    }

    public function getBasename($args)
    {
        $path = $args['path'] ?? '';
        $basename = basename($path);
        return ['basename' => $basename];
    }

    public function formatBytes($args)
    {
        $bytes = $args['bytes'] ?? 0;
        $units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB'];
        
        $i = 0;
        $value = $bytes;
        
        while ($value >= 1024 && $i < count($units) - 1) {
            $value /= 1024;
            $i++;
        }
        
        $formatted = round($value, 2) . ' ' . $units[$i];
        
        return [
            'formatted' => $formatted,
            'value' => round($value, 2),
            'unit' => $units[$i]
        ];
    }
}
