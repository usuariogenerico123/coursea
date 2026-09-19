

export function saveCache(Obj){
    localStorage.setItem("cache", JSON.stringify(Obj));
    return;
}


export function getDataFromCache(){
    const data = localStorage.getItem("cache");
    if(data){
        return JSON.parse(data);
    }
    return null;
}







