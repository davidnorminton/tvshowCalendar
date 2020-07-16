(function() {
    
document.addEventListener('click', function (event) {

    if(event.target.classList.contains('add-show')) {
        addShow(event, event.target.dataset.ref)
    }

    if(event.target.classList.contains('remove-item')) {
        removeShow(event, event.target.dataset.ref)
    }

    if(event.target.id === 'update') {
      updateIcsFile()
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
const updateIcsFile = () => {
    fetch("/updatefile", {mode: "cors"})
    .then(function(response) {
        return response.text()
    }).then(function(){
        if(text === 'updated') {
          console.log("File update");
        } else {
          console.log("error updating file");
        }
    }).catch(function() {
        console.log("error updating file")
    });
};
})()