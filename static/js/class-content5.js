let my_catalogue_id = new URL(location.href).searchParams.get("id")
let div = document.querySelector('div')
let h1 = document.querySelector('h1')
let span = document.querySelector('span')
let button = document.querySelector('button')



async function ask_for_passage() {
    let obj = {
        catalogue_id: my_catalogue_id
    }
    let res = await fetch('http://localhost:8080/book/section', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    h1.innerHTML = result.catalogue_title
    if (result.trialornot == 1) {
        span.style.top = '100px'
        span.style.left = '100px'
        span.style.fontSize = '18px'
        button.style.display = 'none'
        span.innerHTML = result.content
    } else if (result.trialornot == 0 && res2.purchaseornot == 1) {
        span.style.top = '100px'
        span.style.left = '100px'
        button.style.display = 'none'
        span.innerHTML = result.content
        span.style.fontSize = '18px'
    } else if (result.trialornot == 0 && res2.purchaseornot == 0) {
        span.style.top = '250px'
        span.style.left = '320px'
        span.style.fontSize = '26px'
        span.style.color = 'rgb(226, 107, 27)'
        button.style.display = 'none'
        div.style.height='800px'
        span.innerHTML = '您还没有购买本节课程'
    }
}
ask_for_passage()


