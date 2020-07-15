(function() {
    
document.addEventListener('click', function (elem) {
   if(!elem.target.type === 'submit') {
        return
    }
    if(elem.target.classList.contains('add-show')) {
        addShow(elem, elem.target.dataset.ref)
    }
    if(elem.target.classList.contains('remove-item')) {
        removeShow(elem, elem.target.dataset.ref)
    }

}, false);    
// add show to list
const addBtn = (elem) => {
    elem.target.classList.remove('remove-item')
    elem.target.classList.add('add-show')
    elem.target.innerHTML = `ADD SHOW`
};

const addShow = (elem, show) => {
    fetch(`/addshow?add=${show}`, {mode: 'cors'})
    .then(function(response) {
      return response.text();
    })
    .then(function(text) {
      if(text === "error") {
          console.log("error adding show")
      } else {
        removeBtn(elem)
      } 
    })
    .catch(function(error) {
      console.log('Request failed', error)
    });
}; 

// remove show from list
const removeBtn = (elem) => {
    elem.target.classList.remove('add-show')
    elem.target.classList.add('remove-item')
    elem.target.innerHTML = `REMOVE SHOW`
};

const removeShow = (elem, show) => {
    fetch(`/removeshow?remove=${show}`, {mode: 'cors'})
    .then(function(response) {
      return response.text();
    })
    .then(function(text) {
      if(text === "error") {
          console.log("error removing show")
      } else {
        addBtn(elem)
      } 
    })
    .catch(function(error) {
      log('Request failed', error)
    });
};

// update show list

})()