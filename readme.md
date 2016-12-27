##1) Merge sort de archivos
- Archivos con direcciones de correos electrónicos.
- Filtrar correos válidos (expresión regular).
- Ordenar.
- Download archivo ordenado/filtrado.

endpoint: /api/sort

##2) Esteganografía
**a)** Mandar *.bmp y mensaje

- Retornar imagen modificada.

endpoint: /api/bitcode

**b)** Mandar imagen modificada

- Retornar el mensaje original.

endpoint: /api/bitcode/seek

##3) Enviar un grafo no dirigido
- Retornar el árbol abarcador de costo mínimo (Kruskal).

endpoint: /api/kruskal

# Lenguajes
* GO
-run: $ go run server.go
-port: 3000

* NodeJS
-run: $ node server.js
-port: 8000

* Ruby
-run: $ ruby server.rb
-port: 4567

* Python
-run: $ python server.py
-port: 8080

## Desarrollado
Unitec SPS Q5 - 2016