let article_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[0]
let column_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[1]
let boil_data=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[2]

article_data.onclick=function(){
    location.href='/creator/data/content/article'
}
column_data.onclick=function(){
    location.href='/creator/data/content/column'
}
boil_data.onclick=function(){
    location.href='/creator/data/content/pin'
}


let data=document.querySelector('.boil_data_detail').querySelectorAll('h1')


async function ask_for_boil_data(){
    
    let res = await fetch('http://localhost:8080/creator/data/content/pin',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
       
    });
    let res2 = await res.json();
    let result=res2
    data[0].innerHTML=result.pins_number
    data[1].innerHTML=result.pins_like_number 
    data[2].innerHTML=result.pins_comment_number


}
ask_for_boil_data()