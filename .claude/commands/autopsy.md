---
description: Autopsia de un bug ya resuelto — Fidel llena el diagnóstico, tú lo evalúas.
argument-hint: [el bug que acaba de resolverse]
---

# Autopsia del bug

Bug bajo autopsia: $ARGUMENTS

## Filtro previo

Si fue un typo, un slip mecánico o un despiste sin patrón detrás, **dilo y salta la
autopsia**. Convertir cada tropiezo en ceremonia agota el ritual y enseña a temer el
error en vez de leerlo. La autopsia es para bugs que revelaron un modelo mental roto.

## Paso 1 — Fidel llena el informe (bloqueante)

Pídele que complete estos seis campos **él mismo**, antes de que tú digas nada:

- **BUG:** ¿qué pasó?
- **MODELO ORIGINAL:** ¿qué creías que hacía el código / el runtime / la stdlib?
- **MODELO REAL:** ¿qué hace de verdad?
- **SEÑAL PERDIDA:** ¿qué evidencia ya estaba a la vista y podría haberlo revelado antes?
- **CONCEPTO RAÍZ:** ¿qué mecánica de Go se entendió mal?
- **PREVENCIÓN:** ¿qué hábito, test o herramienta lo atrapa la próxima vez?

No los llenes por él. La distinción entre MODELO ORIGINAL y MODELO REAL es la unidad
atómica del aprendizaje: escribir el modelo viejo antes de sustituirlo es lo que hace
que la corrección se fije.

## Paso 2 — Tu evaluación

Corrige lo que esté mal en su diagnóstico, con el *porqué* técnico (memoria, tipos,
runtime, diseño del compilador — según corresponda). Apóyate en `docs/books/` cuando el
concepto esté ahí, y cita capítulo.

Luego clasifica el bug en una de estas categorías:

1. Descuido aislado
2. Patrón repetido
3. Debilidad conceptual
4. Conocimiento de dominio faltante
5. Proceso de depuración débil

La clasificación importa porque cada una pide un remedio distinto: el descuido pide
herramienta (`go vet`, un test), la debilidad conceptual pide lectura, el proceso débil
pide método.

## Paso 3 — Registro

Añade la autopsia condensada (2-4 líneas) a la sección **"Qué costó"** de la entrada
de esta sesión en `JOURNEY.md`, preservando la voz original de Fidel: qué creía, qué
era. No reescribas errores pasados; el registro histórico es el que hace visible el
patrón meses después.

Recomienda un ejercicio de seguimiento **solo** si aumenta su capacidad independiente.
