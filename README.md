"# AlquimIA" 

Forma de ejecucion

main.exe -task {tarea.json}

"#Contenido Task JSON"

Ejemplo

{
    "leng": "golang",
    "version": "1.20.1",
    "info": "considera que el proyecto en golang que haras tiene la configuracion go mod init chatgo,  Se solicita el desarrollo de una aplicación de chat en tiempo real utilizando Golang para el backend y HTML para el frontend, la aplicacion debe tener el poder de exponer el public_html para poder comunicarse con el backend. La aplicación debe ser robusta, escalable y segura, cumpliendo con los estándares modernos de desarrollo de software y desplegada en un entorno de contenedores (Docker) para su fácil implementación en Kubernetes. Requerimientos Funcionales: 1. Registro y Autenticación de Usuarios: - Registro de nuevos usuarios con validación de correo electrónico. - Inicio de sesión seguro con cifrado de contraseñas. - Gestión de sesiones de usuario. 2. Interfaz de Usuario: - Diseño limpio y responsivo utilizando HTML y CSS. - Página principal con lista de contactos y chats recientes. - Ventana de chat con capacidad para enviar y recibir mensajes en tiempo real. - Notificaciones de nuevos mensajes. - Soporte para emojis y archivos adjuntos (imágenes, documentos). 3. Mensajería en Tiempo Real: - Implementación de WebSockets para comunicación en tiempo real. - Envío y recepción de mensajes instantáneamente. - Indicadores de escritura y estado en línea. 4. Almacenamiento de Datos: - Base de datos para almacenar información de usuarios, mensajes y archivos adjuntos. - Utilización de un sistema de base de datos relacional (por ejemplo, PostgreSQL o MySQL). 5. Seguridad: - Cifrado de datos sensibles (contraseñas, mensajes). - Protección contra ataques comunes (XSS, CSRF, SQL Injection). 6. Despliegue: - Configuración de Docker para contenerizar la aplicación. - Archivos de configuración para despliegue en Kubernetes.",
    "source_path": "../task/chatII"
}

