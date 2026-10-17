
import { api } from "./api.js";
// const data = [
//   {
//     "id": 2,
//     "nombre_curso": "Introducción a Go",
//     "nombre_tutor": "Carlos Ma",
//     "video_presentacion": "https://cdn.midominio.com/videos/intro-go.mp4",
//     "metas_aprendizaje": "Aprender los fundamentos del lenguaje Go y su ecosistema",
//     "admin_id": 0,
//     "modulos": [
//       {
//         "id": 1,
//         "titulo_modulo": "Fundamentos del lenguaje",
//         "numero_modulo": 1,
//         "descripcion_modulo": "Variables, tipos de datos, estructuras de control",
//         "curso_id": 2,
//         "temas": [
//           {"id": 1, "numero_tema": "1.1", "titulo_tema": "Variables y constantes", "url_video": "", "duracion": 15, "descripcion": "", "metas_aprendizaje": "Entender la diferencia entre var y :=", "modulo_id": 1},
//           {"id": 2, "numero_tema": "1.2", "titulo_tema": "Variables y constantes", "url_video": "", "duracion": 15, "descripcion": "", "metas_aprendizaje": "Entender la diferencia entre var y :=", "modulo_id": 1},
//           {"id": 4, "numero_tema": "1.3", "titulo_tema": "Variables y constantes", "url_video": "", "duracion": 15, "descripcion": "", "metas_aprendizaje": "Entender la diferencia entre var y :=", "modulo_id": 1}
//         ]
//       },
//       {
//         "id": 2,
//         "titulo_modulo": "Fundamentos del lenguaje",
//         "numero_modulo": 2,
//         "descripcion_modulo": "Variables, tipos de datos, estructuras de control",
//         "curso_id": 2,
//         "temas": [
//           {"id": 6, "numero_tema": "2.1", "titulo_tema": "Variables y constantes", "url_video": "", "duracion": 15, "descripcion": "", "metas_aprendizaje": "Entender la diferencia entre var y :=", "modulo_id": 2}
//         ]
//       }
//     ]
//   }
// ];
console.log("start");

const getData = async ()=>{
    try{
      
       let resp = await api.getAllCourseById(2);
      console.log(resp);
       if(resp.ok == true){
          console.log(resp.data);
          
          return (resp.data);
       }
    }catch(e){
      console.log(e);
      return [];
    }
}


let data = []

//console.log(data.data);


let selectedCourseId = null;
let nextId = 100;

async function renderCourseList() {
  let co = await  getData();
  const list = document.getElementById('courseList');
  if ( co.ok == false){
    console.log(data);
    return
  };
  data.push(co.data);
  console.log(co);
  list.innerHTML =  data.map(c => `
    <div id=${c.id} class="course-item ${c.id === selectedCourseId ? 'active' : ''}" ">
      <span>${c.nombre_curso}</span>
    </div>
  `).join('');
  console.log(data);
  

  document.querySelectorAll(".course-item").forEach(c =>{
      c.addEventListener("click", ()=>{
          console.log(c.id)
          selectedCourseId = c.id;
          //renderCourseList();
          renderMain();
        })
      }
    )
  
}


// MODIFICADO: eliminado código muerto (selectCourse y los bloques comentados que usaban el onclick inline).

function getCourse(id) { 
  console.log("get course by id");
  return data.find(c => c.id == id); 
}

function renderMain() {
  console.log("Rende main");
  const main = document.getElementById('mainContent')
  main.innerHTML='<div class="empty">Selecciona un curso para administrarlo</div>';
  const course = getCourse(selectedCourseId);
  // if (!course) { 
  //   console.log("Not course", course);
  //   main.innerHTML = '<div class="empty">Selecciona un curso para administrarlo</div>'; 
  //   return; 
  // }

  main.innerHTML = `
    <div class="course-header">
      <div class="course-header-info">
        <h2>${course.nombre_curso}</h2>
        <p>Tutor: ${course.nombre_tutor}</p>
        <div class="course-meta">
          <span>${course.modulos.length} módulo(s)</span>
          <span>${course.modulos.reduce((a,m)=>a+m.temas.length,0)} tema(s)</span>
        </div>
      </div>
      <!-- MODIFICADO: onclick inline reemplazado por data-action (patrón de renderCourseList) -->
      <button data-action="add-module" data-course="${course.id}">+ Añadir Módulo</button>
    </div>
    ${course.modulos.length === 0 ? '<div class="empty">No hay módulos. Añade el primero.</div>' :
      course.modulos.map(m => `
        <div class="module">
          <div class="module-header">
            <div class="module-info">
              <h3>Módulo ${m.numero_modulo}: ${m.titulo_modulo}</h3>
              <small>${m.descripcion_modulo}</small>
            </div>
            <div class="module-actions">
              <button class="secondary small" data-action="edit-module" data-course="${course.id}" data-module="${m.id}">Editar</button>
              <button class="danger small" data-action="delete-module" data-course="${course.id}" data-module="${m.id}">Eliminar</button>
              <button class="small" data-action="add-topic" data-course="${course.id}" data-module="${m.id}">+ Tema</button>
            </div>
          </div>
          <div class="topics">
            ${m.temas.length === 0 ? '<small style="color:#86868b">Sin temas</small>' :
              m.temas.map(t => `
                <div class="topic">
                  <div class="topic-info">
                    <strong>${t.numero_tema}. ${t.titulo_tema}</strong>
                    <small>${t.duracion} min</small>
                  </div>
                  <div class="topic-actions">
                    <button class="secondary small" data-action="edit-topic" data-course="${course.id}" data-module="${m.id}" data-topic="${t.id}">Editar</button>
                    <button class="danger small" data-action="delete-topic" data-course="${course.id}" data-module="${m.id}" data-topic="${t.id}">Eliminar</button>
                  </div>
                </div>
              `).join('')
            }
          </div>
        </div>
      `).join('')
    }
  `;

  // MODIFICADO: listeners enganchados tras pintar el HTML (patrón de renderCourseList).
  // Los ids de dataset llegan como string, por eso se convierten con Number()
  // para que funcionen los === de addModule/editModule/deleteModule/addTopic/editTopic/deleteTopic.
  main.querySelectorAll("[data-action]").forEach(btn => {
    btn.addEventListener("click", () => {
      const { action, course, module: moduleId, topic } = btn.dataset;
      switch (action) {
        case "add-module":    addModule(Number(course)); break;
        case "edit-module":   editModule(Number(course), Number(moduleId)); break;
        case "delete-module": deleteModule(Number(course), Number(moduleId)); break;
        case "add-topic":     addTopic(Number(course), Number(moduleId)); break;
        case "edit-topic":    editTopic(Number(course), Number(moduleId), Number(topic)); break;
        case "delete-topic":  deleteTopic(Number(course), Number(moduleId), Number(topic)); break;
      }
    });
  });
}

function openModal(title, body, onSave) {
  document.getElementById('modalTitle').textContent = title;
  document.getElementById('modalBody').innerHTML = body;
  document.getElementById('modal').classList.add('open');
  document.getElementById('modalSave').onclick = onSave;
}

document.querySelector(".closeModal").addEventListener("click", closeModal)
function closeModal() { document.getElementById('modal').classList.remove('open'); }


document.querySelector(".addCourse").addEventListener("click", addCourse)
function addCourse() {
  openModal('Nuevo Curso', `
    <div class="form-row"><label>Nombre del curso</label><input id="f_nombre" type="text"></div>
    <div class="form-row"><label>Tutor</label><input id="f_tutor" type="text"></div>
    <div class="form-row"><label>Video presentación (URL)</label><input id="f_video" type="text"></div>
    <div class="form-row"><label>Metas de aprendizaje</label><textarea id="f_metas"></textarea></div>
  `, () => {
    const nuevo = {
      id: nextId++,
      nombre_curso: document.getElementById('f_nombre').value || 'Sin nombre',
      nombre_tutor: document.getElementById('f_tutor').value || '',
      video_presentacion: document.getElementById('f_video').value || '',
      metas_aprendizaje: document.getElementById('f_metas').value || '',
      admin_id: 0,
      modulos: []
    };
    data.push(nuevo);
    selectedCourseId = nuevo.id;
    closeModal();
    renderCourseList();
    renderMain();
  });
}

function addModule(courseId) {
  const course = getCourse(courseId);
  openModal('Nuevo Módulo', `
    <div class="form-row"><label>Número de módulo</label><input id="f_num" type="number" value="${course.modulos.length + 1}"></div>
    <div class="form-row"><label>Título</label><input id="f_titulo" type="text"></div>
    <div class="form-row"><label>Descripción</label><textarea id="f_desc"></textarea></div>
  `, () => {
    course.modulos.push({
      
      id: nextId++,
      numero_modulo: parseInt(document.getElementById('f_num').value) || 1,
      titulo_modulo: document.getElementById('f_titulo').value || 'Sin título',
      descripcion_modulo: document.getElementById('f_desc').value || '',
      curso_id: courseId,
      temas: []
    });
    closeModal();
    renderMain();
  });
}

function editModule(courseId, moduleId) {
  const course = getCourse(courseId);
  const m = course.modulos.find(x => x.id === moduleId);
  openModal('Editar Módulo', `
    <div class="form-row"><label>Número de módulo</label><input id="f_num" type="number" value="${m.numero_modulo}"></div>
    <div class="form-row"><label>Título</label><input id="f_titulo" type="text" value="${m.titulo_modulo}"></div>
    <div class="form-row"><label>Descripción</label><textarea id="f_desc">${m.descripcion_modulo}</textarea></div>
  `, () => {
    m.numero_modulo = parseInt(document.getElementById('f_num').value) || m.numero_modulo;
    m.titulo_modulo = document.getElementById('f_titulo').value || m.titulo_modulo;
    m.descripcion_modulo = document.getElementById('f_desc').value;
    closeModal();
    renderMain();
  });
}

function deleteModule(courseId, moduleId) {
  console.log("deleteModule", "courseId:", courseId, typeof courseId, "moduleId:", moduleId, typeof moduleId);
  if (!confirm('¿Eliminar este módulo y todos sus temas?')) return;
  const course = getCourse(courseId);
  course.modulos = course.modulos.filter(m => m.id !== moduleId);
  renderMain();
}

function addTopic(courseId, moduleId) {
  const course = getCourse(courseId);
  const m = course.modulos.find(x => x.id === moduleId);
  openModal('Nuevo Tema', `
    <div class="form-row"><label>Número de tema</label><input id="f_num" type="text" value="${m.numero_modulo}.${m.temas.length + 1}"></div>
    <div class="form-row"><label>Título</label><input id="f_titulo" type="text"></div>
    <div class="form-row"><label>Duración (min)</label><input id="f_dur" type="number" value="15"></div>
    <div class="form-row"><label>URL del video</label><input id="f_url" type="text"></div>
    <div class="form-row"><label>Descripción</label><textarea id="f_desc"></textarea></div>
    <div class="form-row"><label>Metas de aprendizaje</label><textarea id="f_metas"></textarea></div>
  `, () => {
    m.temas.push({
      id: nextId++,
      numero_tema: document.getElementById('f_num').value || `${m.numero_modulo}.${m.temas.length+1}`,
      titulo_tema: document.getElementById('f_titulo').value || 'Sin título',
      url_video: document.getElementById('f_url').value || '',
      duracion: parseInt(document.getElementById('f_dur').value) || 0,
      descripcion: document.getElementById('f_desc').value || '',
      metas_aprendizaje: document.getElementById('f_metas').value || '',
      modulo_id: moduleId
    });
    closeModal();
    renderMain();
  });
}

function editTopic(courseId, moduleId, topicId) {
  const course = getCourse(courseId);
  const m = course.modulos.find(x => x.id === moduleId);
  const t = m.temas.find(x => x.id === topicId);
  openModal('Editar Tema', `
    <div class="form-row"><label>Número de tema</label><input id="f_num" type="text" value="${t.numero_tema}"></div>
    <div class="form-row"><label>Título</label><input id="f_titulo" type="text" value="${t.titulo_tema}"></div>
    <div class="form-row"><label>Duración (min)</label><input id="f_dur" type="number" value="${t.duracion}"></div>
    <div class="form-row"><label>URL del video</label><input id="f_url" type="text" value="${t.url_video}"></div>
    <div class="form-row"><label>Descripción</label><textarea id="f_desc">${t.descripcion}</textarea></div>
    <div class="form-row"><label>Metas de aprendizaje</label><textarea id="f_metas">${t.metas_aprendizaje}</textarea></div>
  `, () => {
    t.numero_tema = document.getElementById('f_num').value;
    t.titulo_tema = document.getElementById('f_titulo').value;
    t.duracion = parseInt(document.getElementById('f_dur').value) || 0;
    t.url_video = document.getElementById('f_url').value;
    t.descripcion = document.getElementById('f_desc').value;
    t.metas_aprendizaje = document.getElementById('f_metas').value;
    closeModal();
    renderMain();
  });
}

function deleteTopic(courseId, moduleId, topicId) {
  console.log("deleteTopic", "courseId:", courseId, typeof courseId, "moduleId:", moduleId, typeof moduleId, "topicId:", topicId, typeof topicId);
  if (!confirm('¿Eliminar este tema?')) return;
  const course = getCourse(courseId);
  const m = course.modulos.find(x => x.id === moduleId);
  m.temas = m.temas.filter(t => t.id !== topicId);
  renderMain();
}


renderCourseList()