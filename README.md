# Desarrollo de una API REST para gestión de productos

La empresa necesita una API REST que permita la gestión de productos en un sistema de e-commerce. La API debe permitir crear, leer, actualizar y eliminar productos. Los productos tienen atributos como nombre, precio, stock y categoría. La API debe validar que los nombres de los productos no sean duplicados y que los precios no sean negativos. En caso de que una validación falle, la API debe devolver un mensaje de error descriptivo. La API debe ser idempotente para las operaciones de creación y actualización de productos.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST con Go, Gin framework y GORM |
| **Nivel** | junior-l1 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Creación de la estructura básica de la API

**Objetivo:** Tener una API que permita crear productos con validación de nombre y precio

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe permitir crear productos con los atributos nombre, precio, stock y categoría.
- La API debe validar que el nombre del producto no sea duplicado y que el precio no sea negativo.
- La API debe ser idempotente para la operación de creación de productos.

**Entregable:** API que permite crear productos con validación de nombre y precio

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo estructurar los datos para asegurar la unicidad del nombre y la validez del precio.
- Piensa en cómo implementar la idempotencia para la creación de productos.

</details>

### Fase 2: Implementación de la lectura y actualización de productos

**Objetivo:** Ampliar la API para permitir leer y actualizar productos

**Tiempo estimado:** 3 horas

**Instrucciones:**

- La API debe permitir leer productos por su ID.
- La API debe permitir actualizar productos con validación de nombre y precio.
- La API debe ser idempotente para la operación de actualización de productos.

**Entregable:** API que permite crear, leer y actualizar productos con validación de nombre y precio

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo implementar la lectura de productos por ID.
- Piensa en cómo asegurar la idempotencia para la actualización de productos.

</details>

### Fase 3: Implementación de la eliminación de productos

**Objetivo:** Completar la API para permitir eliminar productos

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe permitir eliminar productos por su ID.

**Entregable:** API completa que permite crear, leer, actualizar y eliminar productos con validación de nombre y precio

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo implementar la eliminación de productos por ID.

</details>

### Fase 4: Refactorización y optimización de la API

**Objetivo:** Optimizar y refactorizar la API para mejorar su rendimiento y mantenibilidad

**Tiempo estimado:** 1 hora

**Instrucciones:**

- Identifica y soluciona posibles ineficiencias en la API.
- Refactoriza el código para mejorar su legibilidad y mantenibilidad.

**Entregable:** API refactorizada y optimizada

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo identificar y solucionar ineficiencias en la API.
- Piensa en cómo refactorizar el código para mejorar su legibilidad y mantenibilidad.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una API REST y cómo se utiliza en este reto?
- **paraQueSirve**: ¿Para qué sirve la validación de nombres y precios en la API?
- **comoSeUsa**: ¿Cómo se usa la idempotencia en las operaciones de creación y actualización de productos?
- **erroresComunes**: ¿Qué errores comunes puedes encontrar al implementar una API REST?
- **queDecisionesImplica**: ¿Qué decisiones debes tomar al refactorizar y optimizar la API?

## Criterios de Evaluacion

- Implementación correcta de la API con validación de nombres y precios.
- Idempotencia en las operaciones de creación y actualización de productos.
- Optimización y refactorización del código para mejorar rendimiento y mantenibilidad.

## Como trabajar con un asistente de IA

- **AGENTS.md** — instrucciones nativas del repo (Cursor, Codex, Copilot, Gemini, Claude Code). Abrí el proyecto y el agente las carga solo.
- **PROMPT_MEJORA.md** — el mismo prompt, para copiar y pegar en un chat (claude.ai, ChatGPT, etc.).

---

*Reto generado automaticamente por Challenge Generator - Pragma*
