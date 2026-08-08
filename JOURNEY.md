# Bitácora del camino

## 2026-08-07 — Sesión 1: primera CLI (Learning Go caps. 1-2)
- **Qué se hizo:** módulo `01-fundamentos` con `filereader.go` — CLI que valida argumentos, lee un archivo completo con `os.ReadFile` y maneja errores con mensaje a stderr + exit code correcto. Las 3 pruebas (archivo válido, sin args, archivo inexistente) pasan; `go vet` y `gofmt` limpios.
- **Qué se aprendió:** `os.Args[0]` es el nombre del programa (el bug del uróboro: el programa se leyó a sí mismo y escupió un header ELF); slice `nil` ≠ slice corto → validar con `len`; error-valor vs. panic del runtime (los panics se previenen, no se cachean); `panic`/`check` de Go by Example es anti-patrón — el idioma es `if err != nil` en el sitio; `os.Exit` no ejecuta los `defer`; errores a stderr, no stdout; verbos de `fmt` eligen representación (`%T`, `%v`, `%q`, `%s`), no declaran tipo como en C; `os.Open` (bajo nivel, trozos) vs. `os.ReadFile` (alto nivel, todo en memoria).
- **Qué costó:** la condición de longitud de `os.Args` (dos intentos fallidos: `== nil`, `< 1`); descifrar el output binario del primer run; soltar el hábito de copiar patrones sin entenderlos.
- **Qué sigue:** commit + push (crear repo en GitHub sigue pendiente). Semana 2: Learning Go cap. 3 (strings, slices, maps) → tokenizador: partir texto en palabras y contar frecuencias.

## 2026-08-07 — Día 0: preparación
- Proyecto creado, reglas de mentoría definidas en CLAUDE.md, repo inicializado.
- Biblioteca montada: los 4 libros (.epub) convertidos a markdown por capítulo en `docs/books/`, con `INDEX.md`. Los libros quedan solo locales (gitignore) por copyright.
- Decisiones: 3-5 h/semana reales; aprendizaje guiado por proyecto desde el día 1; proyecto norte = motor de búsqueda full-text en Go (fuera del dominio de MELI, sin riesgo de IP); CLRS solo como referencia just-in-time, no como track paralelo; check-in semanal automatizado.
- Plan de estudio completo en `PLAN.md`.
- Pendiente: crear repo en GitHub y hacer push; primera sesión de código (Learning Go cap. 1-2 → CLI mínima).
