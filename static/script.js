// static/script.js
$(document).ready(function() {
    $("#uploadForm").submit(function(e) {
        e.preventDefault();

        var formData = new FormData(this);

        $.ajax({
            type: "POST",
            url: "/upload",
            data: formData,
            contentType: false,
            processData: false,
            success: function(response) {
                showSuccess(response.message, response.textResult);
            },
            error: function(xhr, status, error) {
                showError(xhr.responseText);
            }
        });
    });

    function showSuccess(message, textResult) {
        $("#resultContainer").show();
        $("#successMessage").html("<strong>Success!</strong> " + message);
        $("#textResult").html("<strong>Text Result:</strong> " + textResult);
        $("#errorMessageContainer").hide();
    }

    function showError(errorMessage) {
        $("#errorMessageContainer").show();
        $("#errorMessage").html("<strong>Error!</strong> " + errorMessage);
        $("#resultContainer").hide();
    }
});
