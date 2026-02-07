<?php

namespace App\Libraries\Tools;

class MathTools
{
    public function getTools()
    {
        return [
            [
                'name' => 'math_add',
                'description' => 'Penjumlahan dua angka',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'a' => [
                            'type' => 'number',
                            'description' => 'Angka pertama'
                        ],
                        'b' => [
                            'type' => 'number',
                            'description' => 'Angka kedua'
                        ]
                    ],
                    'required' => ['a', 'b']
                ],
                'handler' => [$this, 'add']
            ],
            [
                'name' => 'math_subtract',
                'description' => 'Pengurangan dua angka',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'a' => [
                            'type' => 'number',
                            'description' => 'Angka pertama'
                        ],
                        'b' => [
                            'type' => 'number',
                            'description' => 'Angka kedua'
                        ]
                    ],
                    'required' => ['a', 'b']
                ],
                'handler' => [$this, 'subtract']
            ],
            [
                'name' => 'math_multiply',
                'description' => 'Perkalian dua angka',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'a' => [
                            'type' => 'number',
                            'description' => 'Angka pertama'
                        ],
                        'b' => [
                            'type' => 'number',
                            'description' => 'Angka kedua'
                        ]
                    ],
                    'required' => ['a', 'b']
                ],
                'handler' => [$this, 'multiply']
            ],
            [
                'name' => 'math_divide',
                'description' => 'Pembagian dua angka',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'a' => [
                            'type' => 'number',
                            'description' => 'Angka pembilang'
                        ],
                        'b' => [
                            'type' => 'number',
                            'description' => 'Angka penyebut'
                        ]
                    ],
                    'required' => ['a', 'b']
                ],
                'handler' => [$this, 'divide']
            ],
            [
                'name' => 'math_power',
                'description' => 'Pangkat angka (a^b)',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'base' => [
                            'type' => 'number',
                            'description' => 'Angka dasar'
                        ],
                        'exponent' => [
                            'type' => 'number',
                            'description' => 'Pangkat'
                        ]
                    ],
                    'required' => ['base', 'exponent']
                ],
                'handler' => [$this, 'power']
            ],
            [
                'name' => 'math_sqrt',
                'description' => 'Akar kuadrat dari angka',
                'parameters' => [
                    'type' => 'object',
                    'properties' => [
                        'number' => [
                            'type' => 'number',
                            'description' => 'Angka yang akan dihitung akar kuadratnya'
                        ]
                    ],
                    'required' => ['number']
                ],
                'handler' => [$this, 'sqrt']
            ]
        ];
    }

    public function add($args)
    {
        $a = $args['a'] ?? 0;
        $b = $args['b'] ?? 0;
        $result = $a + $b;
        return ['result' => $result];
    }

    public function subtract($args)
    {
        $a = $args['a'] ?? 0;
        $b = $args['b'] ?? 0;
        $result = $a - $b;
        return ['result' => $result];
    }

    public function multiply($args)
    {
        $a = $args['a'] ?? 0;
        $b = $args['b'] ?? 0;
        $result = $a * $b;
        return ['result' => $result];
    }

    public function divide($args)
    {
        $a = $args['a'] ?? 0;
        $b = $args['b'] ?? 0;
        
        if ($b == 0) {
            throw new \Exception('Division by zero');
        }
        
        $result = $a / $b;
        return ['result' => $result];
    }

    public function power($args)
    {
        $base = $args['base'] ?? 0;
        $exponent = $args['exponent'] ?? 0;
        $result = pow($base, $exponent);
        return ['result' => $result];
    }

    public function sqrt($args)
    {
        $number = $args['number'] ?? 0;
        
        if ($number < 0) {
            throw new \Exception('Cannot calculate square root of negative number');
        }
        
        $result = sqrt($number);
        return ['result' => $result];
    }
}
