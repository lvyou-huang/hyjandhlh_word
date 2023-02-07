let list = document.querySelector('.list')
let myid = new URL(location.href).searchParams.get("id")
async function ask_for_column() {
    let obj = {
        user_id:myid
    }
    let res = await fetch('http://localhost:8080/getcolumns', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result=res2.msg;
    for(let i=0;i<result.length;i++){
        let div=document.createElement('div');
        list.appendChild(div)
        let img=document.createElement('img')
        let h2=document.createElement('h2')
        let span1=document.createElement('span')
        let span2=document.createElement('span')

        div.appendChild(img)
        div.appendChild(h2)
        div.appendChild(span1)
        div.appendChild(span2)

    

        img.src=result[i].cover
        h2.innerHTML=result[i].column_title
        span1.innerHTML=result[i].column_intruduction
        span2.innerHTML=result[i].time
       img.id=result[i].id
        h2.id=result[i].id
        span1.id=result[i].id
        span2.id=result[i].id
        div.id=result[i].id
        div.onclick=function(e){
            location.href='/column?id='+e.target.id
        }
    }
}
ask_for_column()