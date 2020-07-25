(function() {
    
const msg = document.getElementById('messages');


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

const clearMsg = () => {
    msg.classList.remove('error-msg');
    msg.classList.remove('good-msg');
    msg.classList.remove('amber-msg');
    msg.innerHTML = '';
};

const addShow = (elem, show) => {
    clearMsg()
    fetch(`/addshow?add=${show}`, {mode: 'cors'})
    .then(function(response) {
      return response.text();
    })
    .then(function(text) {
      if(text === "error") {
          msg.innerHTML = `There was an error removing ${show} to your list`
          msg. classList.add('error-msg');       
      } else {
        removeBtn(elem)
        msg.innerHTML = `${show} was added to your Calendar file`;
        msg.classList.add('good-msg');        
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
    clearMsg();
    fetch(`/removeshow?remove=${show}`, {mode: 'cors'})
    .then(function(response) {
      return response.text();
    })
    .then(function(text) {
      if(text === "error") {
          msg.innerHTML = `There was an error adding ${show} to your list`
          msg. classList.add('error-msg');  
      } else {
          addBtn(elem)
          msg.innerHTML = `${show} was removed to your Calendar file`;
          msg.classList.add('good-msg');
      } 
    })
    .catch(function(error) {
      log('Request failed', error)
    });
};


const addSpinner = () => {
    msg.innerHTML = `<i class="icon-spinner spinner"></i>`
    msg.classList.add("amber-msg")
};

// update show list
const updateIcsFile = () => {
    clearMsg()
    addSpinner()
    fetch("/updatefile", {mode: "cors"})
    .then(function(response) {
        return response.text()
    }).then(function(text){
        clearMsg();
        if(text === 'updated') {
            msg.innerHTML = "Your Calendar file has been updated";
            msg.classList.add('good-msg');
        } else {
            msg.innerHTML = "Failed to update your calendar file"
            msg. classList.add('error-msg')
        }
    }).catch(function() {
        console.log("error updating file")
    });
};
})()

