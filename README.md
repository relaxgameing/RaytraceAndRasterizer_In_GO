# Objective
the main reason of this project was to start my journey in the feild of computer graphics. 
this repo is my implementation of the book [Computer Graphics from Scratch by Gabriel Gambetta](https://gabrielgambetta.com/computer-graphics-from-scratch/).

## Feature / Concepts implemented
1. RayTracer
   * Lighting - Ambient , Directional , Point
   * Shadows
   * Reflections - Diffuse , Specular
2. Rasterizer
   * Drawing Primitive - Lines , Triangle
   * Filling and Shading Triangles
   * Perspective Projection
   * Custom `.obj` file parser for rendering a model
   * Rendering scenes from `.json` files
   * View-Frustum Culling [partial implementation]
   * BackFace Culling with Depth Buffer

# RayTracing
<img width="912" height="744" alt="Screenshot 2026-03-12 at 6 41 36 PM" src="https://github.com/user-attachments/assets/fa625dde-d825-4c02-8606-e18977abeff8" />

this is the output of my implementation of the raytracer with depth = 4 

# Rasterizer
https://github.com/user-attachments/assets/07782af3-a450-4ceb-a9da-d3267bcfd55a

this is the basic rasterizer without lighting 

# Usage
To run `Rasterizer`
```shell 
make rs
````

To run `Raytracer`
```shell
make ray
```

### To define your own scene in `Rasterizer` 
1. create necessary models that will be used in the scene inside `./scene` folder
2. create a `.json` file which would describe your scene ( refer the `./scene/rasterizaor.json` for structure)
3. update the `func NewRasterizationRequirements() *OptionRequirement` in `config.go` file, to provide the relative path to your scene

### To define your own scene in `Raytracer`
1. Create your primitive or model which satisfies the `Entity` Interface
2. In `func NewRayTracingRequirements() *OptionRequirement` in `config.go` file and add the entities in the scene like
   ```go
   scene.AddSceneEntities(
  		entity.NewMyEntity(
  			<your entity paramters>
     ),
   )
   ```
