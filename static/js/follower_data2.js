let changing0=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[0]
let changing1=document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[1]
changing0.onclick=function(){
    location.href='/creator/data/follow/data'
}
changing1.onclick=function(){
    location.href='/creator/data/follower/list'
}


let data=document.querySelector('.follower_data_detail').querySelectorAll('h1')

let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let times = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;
async function ask_for_follower_data(){ 
    let obj={
        time:times
    }
    let res = await fetch('http://localhost:8080/creator/data/follow/data',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result=res2
    data[0].innerHTML=result.follower_number
    data[1].innerHTML=result.new_follower_number  
    data[2].innerHTML=result.cancel_number
    data[3].innerHTML=result.added_number
}
ask_for_follower_data()