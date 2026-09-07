# Prompt para Mejorar el Codigo Base

Copia y pega el siguiente contenido completo en un asistente de IA (Claude, ChatGPT, etc.)
para obtener un ZIP con el proyecto arrancable. Si el adjunto es una carcasa (docs/placeholders),
el asistente debe materializar la estructura del stack del briefing, sin resolver las fases del reto.

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
API REST con Go, Gin framework y GORM

### Reto
- Tema: API REST con Go, Gin framework y GORM
- Seniority: junior-l1
- Tipo: practical
- Título: Desarrollo de una API REST para gestión de productos
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Creación de la estructura básica de la API — objetivo: Tener una API que permita crear productos con validación de nombre y precio — entregable (NO resolver): API que permite crear productos con validación de nombre y precio
- Fase 2: Implementación de la lectura y actualización de productos — objetivo: Ampliar la API para permitir leer y actualizar productos — entregable (NO resolver): API que permite crear, leer y actualizar productos con validación de nombre y precio
- Fase 3: Implementación de la eliminación de productos — objetivo: Completar la API para permitir eliminar productos — entregable (NO resolver): API completa que permite crear, leer, actualizar y eliminar productos con validación de nombre y precio
- Fase 4: Refactorización y optimización de la API — objetivo: Optimizar y refactorizar la API para mejorar su rendimiento y mantenibilidad — entregable (NO resolver): API refactorizada y optimizada

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"src/controllers"
	"src/models"
	"src/services"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
		panic("failed to connect database")
	}

	productService := services.NewProductService(db)
	productController := controllers.NewProductController(productService)

	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetProduct)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	r.Run(":8080")
}
// === ARCHIVO: src/go.mod ===
module github.com/yourusername/yourmodule

go 1.18

require (
	github.com/gin-gonic/gin v1.7.7
	gorm.io/gorm v1.21.12
)
// === ARCHIVO: src/controllers/product_controller.go ===
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/services"
	"src/utils"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var productDTO services.CreateProductDTO
	if err := c.ShouldBindJSON(&productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateProductDTO(productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.productService.CreateProduct(productDTO)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.productService.GetProduct(id)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var productDTO services.UpdateProductDTO
	if err := c.ShouldBindJSON(&productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := utils.ValidateProductDTO(productDTO); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := pc.productService.UpdateProduct(id, productDTO)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := pc.productService.DeleteProduct(id); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
// === ARCHIVO: src/models/product.go ===
package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `gorm:"unique;not null"`
	Price   float64 `gorm:"not null"`
	Stock   int     `gorm:"not null"`
	Category string  `gorm:"not null"`
}
// === ARCHIVO: src/services/product_service.go ===
package services

import (
	"errors"
	"gorm.io/gorm"
	"src/models"
)

type CreateProductDTO struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Category string  `json:"category"`
}

type UpdateProductDTO struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Category string  `json:"category"`
}

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (ps *ProductService) CreateProduct(productDTO CreateProductDTO) (models.Product, error) {
	product := models.Product{
		Name:    productDTO.Name,
		Price:   productDTO.Price,
		Stock:   productDTO.Stock,
		Category: productDTO.Category,
	}
	if err := ps.db.Create(&product).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) GetProduct(id string) (models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(id string, productDTO UpdateProductDTO) (models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return models.Product{}, err
	}
	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Stock = productDTO.Stock
	product.Category = productDTO.Category
	if err := ps.db.Save(&product).Error; err!= nil {
		return models.Product{}, err
	}
	return product, nil
}

func (ps *ProductService) DeleteProduct(id string) error {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err!= nil {
		return err
	}
	if err := ps.db.Delete(&product).Error; err!= nil {
		return err
	}
	return nil
}
// === ARCHIVO: src/utils/validation.go ===
package utils

import (
	"errors"
	"src/services"
)

func ValidateProductDTO(productDTO services.CreateProductDTO) error {
	if productDTO.Name == "" {
		return errors.New("name is required")
	}
	if productDTO.Price < 0 {
		return errors.New("price must be positive")
	}
	return nil
}

```
