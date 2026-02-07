<?php

namespace App\Libraries\Tools;

class StringTools
{
    public function getTools()
    {
        return [
            [
                'name' => 'string_reverse',
                'description' => 'Membalikkan urutan karakter dalam string',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'text' => [
                            'type' => 'string',
                            'description' => 'Text yang akan dibalik'
                        ]
                    ],
                    'required' => ['text']
                ],
                'handler' => [$this, 'reverse']
            ],
            [
                'name' => 'string_uppercase',
                'description' => 'Ubah string menjadi huruf besar (uppercase)',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'text' => [
                            'type' => 'string',
                            'description' => 'Text yang akan diubah'
                        ]
                    ],
                    'required' => ['text']
                ],
                'handler' => [$this, 'uppercase']
            ],
            [
                'name' => 'string_lowercase',
                'description' => 'Ubah string menjadi huruf kecil (lowercase)',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'text' => [
                            'type' => 'string',
                            'description' => 'Text yang akan diubah'
                        ]
                    ],
                    'required' => ['text']
                ],
                'handler' => [$this, 'lowercase']
            ],
            [
                'name' => 'string_length',
                'description' => 'Hitung panjang string',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'text' => [
                            'type' => 'string',
                            'description' => 'Text yang akan dihitung panjangnya'
                        ]
                    ],
                    'required' => ['text']
                ],
                'handler' => [$this, 'length']
            ],
            [
                'name' => 'string_trim',
                'description' => 'Hapus whitespace di awal dan akhir string',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'text' => [
                            'type' => 'string',
                            'description' => 'Text yang akan di-trim'
                        ]
                    ],
                    'required' => ['text']
                ],
                'handler' => [$this, 'trim']
            ]
        ];
    }

    public function reverse($args)
    {
        $text = $args['text'] ?? '';
        $result = strrev($text);
        return ['reversed' => $result];
    }

    public function uppercase($args)
    {
        $text = $args['text'] ?? '';
        $result = strtoupper($text);
        return ['uppercase' => $result];
    }

    public function lowercase($args)
    {
        $text = $args['text'] ?? '';
        $result = strtolower($text);
        return ['lowercase' => $result];
    }

    public function length($args)
    {
        $text = $args['text'] ?? '';
        $length = strlen($text);
        return ['length' => $length];
    }

    public function trim($args)
    {
        $text = $args['text'] ?? '';
        $result = trim($text);
        return ['trimmed' => $result];
    }
}
