# SegInf Challenge

La herramienta desarrollada permite la creación y consulta (mediante llamados tipo  **GET** y  **POST**) de **Items** y **Files**. La misma esta compuesta por 3 containers de  **Docker** que se comunican entre si. Los containers ejecutan cada uno las siguientes herramientas

- **Nginx**      : Se encarga de distribuir el trafico recibido a la rest app
- **Go rest app**: Se encarga de la logica de crear los objetos 
- **Mysql**      : Se utiliza para almacenar los Items y los Id de los Files creados en Drive

# Items
Los Items son objetos que se pueden crear mediante un POST a la app. El lamado debe constar con los siguientes parametros:

- ID
- Name
- Descripción

Los Items almacenados en una base MySql de manera local. Se pueden consultar los objetos creados mediante llamados de tipo GET

# Files
Al igual que los Items, son objetos que se pueden crear mediante un POST a la app. El lamado debe constar con los siguientes parametros:
- Name
- Descripcion

Los Files son almacenados en la nube, utilizando el servicio que brinda Google Drive. La herramienta almacena localmente el Id que posee el archivo en drive. (El mismo se obtiene al crear el archivo en Drive)

Dado un Id de un archivo creado en Drive, se puede consultar si el mismo contiene en su descripcion una palabra pasada por parametro

# Problemas

Los principales problemas con los que me encontre fueron los siguientes:

- Falta de experiencia con GoLang
- Aprender sobre Docker
- Autenticacion de Drive

La solución a los primeros dos problemas fue ponerme a investigar e ir probando varias soluciones que se me ocurian. Cuando quedaba conforme tanto con su funcionamiento como con la calidad de la solucion implementada, seguia adelante.

La autenticacion en Drive fue otro problema. Cuando llega el momento de hacer un POST o un GET a la api de Drive, es necesario autenticar la aplicacion, para esto hay que pasarle a un metodo de la libreria de Drive un token. El mismo debe ser  copiado de una pagina y pegado por STD-IN pero luego de printear por STD-OUT la url de donde conseguir el token, la aplicacion crasheaba instantaneamente. La solucion fue modificar ligeramente la libreria de Drive para que buscara el token en una variable