function prompt() {
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end"
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }
    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    async function custom(c) {
        const {
            icon = "",
            msg = "",
            title = "",
            showConfirmButton = true,
        } = c;


        const { value: result } = await Swal.fire({
            icon: icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton: showConfirmButton,
            willOpen: () => {
                if (c.willOpen  !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel){
                if (result.value !== ""){
                    if (c.callback !== undefined){
                        c.callback(result)
                    }
                }else{
                    c.callback(false)
                }
            } else {
                c.callback(false)
            }
        }

    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}


function booking(csrf_token, room_id) {
    let attention = prompt();
    // notify("This is a test", "success")
    //notifyModel("title", "<em>Hello world!</em>", "success", "My button Text")
    let html = `

    <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
        <div class="d-flex flex-row justify-content-evenly" id="reservation-dates-modal">
            <div class="p-2">
                <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival">
            </div>
            <div class="p-2">
                <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure">
            </div>
        </div>
    </form>
    `
    attention.custom({
        msg: html,
         title: "Choose your dates",
         willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                    showOnFocus: true,
                    orientation: "top",
                    minDate: new Date(),
                })

            },
         didOpen: () => {
                document.getElementById('start').removeAttribute('disabled')
                document.getElementById('end').removeAttribute('disabled')
            },
         callback: function(result) {
            console.log("called");
            let form = document.getElementById("check-availability-form")
            let formData = new FormData(form)
            formData.append("csrf_token", csrf_token);
            formData.append("room_id", room_id);

            fetch('/search-availability-json', {
               method: "post",
               body:formData,
        })
            .then(response => response.json())
            .then(data => {
                if (data.ok) {
                    attention.custom({
                    icon: "success",
                    msg: '<p> Room is available!</p>'
                        + '<p><a href="/book-room?id='
                        + data.room_id
                        + '&s='
                        + data.start_date
                        + '&e='
                        + data.end_date
                        + '" class="btn btn-primary">'
                        + 'Book Now!</a></p>',
                    showConfirmButton: false
                    })
                }else{
                    attention.error({
                            msg: "No availability",
                        })
                    }
                })

             }
        });
    }