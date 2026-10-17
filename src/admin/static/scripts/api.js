


const API_BASE_PATH = "https://5097273d78f12663-135-237-130-232.serveousercontent.com" + "/api/course"; 

export const api = {

    request : async (path, options = {}, expectedResponse)=>{

                try {
                    const res = await fetch(`${API_BASE_PATH}${path}`, {
                        headers: { "Content-Type": "application/json", ...(options.headers || {}) },
                        ...options,
                    })
                    const data = await res.json();
                    if (res.status === expectedResponse) {
                        return { "ok": true, "data": data }
                    }
                    return { "ok": false, "data": data }
                } catch (e) {
                    return { "ok": false, "data": e }
                }
    },

    getAllCourseById: (id)=>{
        return api.request(`/${id}`, { method: "GET" }, 200)
    },

    saveCourse:(courseData)=>{
        return api.request(`/`, { method: "POST", body: JSON.stringify(courseData) }, 201)
    },

    saveModule:(idCourse, moduleData)=>{
        return api.request(`/module/${idCourse}`, { method: "POST", body: JSON.stringify(moduleData) }, 201)
    },

    saveTheme:(idModule, themeData)=>{
        return api.request(`/module/theme/${idModule}`, { method: "POST", body: JSON.stringify(themeData) }, 201)
    },

    updateCourse:(id, courseData)=>{
        return api.request(`/${id}`, { method: "PUT", body: JSON.stringify(courseData) }, 200)
    },

    updateModule:(id, moduleData)=>{
        return api.request(`/module/${id}`, { method: "PUT", body: JSON.stringify(moduleData) }, 200)
    },

    updateTheme:(id, themeData)=>{
        return api.request(`/module/theme/${id}`, { method: "PUT", body: JSON.stringify(themeData) }, 200)
    },

    deleteCourse:(id)=>{
        return api.request(`/${id}`, { method: "DELETE" }, 204)
    },

    deleteModule:(id)=>{
        return api.request(`/module/${id}`, { method: "DELETE" }, 204)
    },

    deleteTheme:(id)=>{
        return api.request(`/module/theme/${id}`, { method: "DELETE" }, 204)
    }
}
