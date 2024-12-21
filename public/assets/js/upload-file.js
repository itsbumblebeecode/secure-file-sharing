// Disable Dropzone's automatic discovery of elements
Dropzone.autoDiscover = false;

// Initialize a Dropzone instance for the element with ID "sfs-dropzone"
var sfsDropzone = new Dropzone("#sfs-dropzone", {
    url: "/api/upload",      // API endpoint for file uploads
    paramName: "file",       // Parameter name for the uploaded file in the form data
    maxFilesize: 2,          // Maximum file size allowed in MB
});

// Event listener for the "success" event, triggered when a file is successfully uploaded
sfsDropzone.on("success", function(file, response) {
    // Validate the response to ensure a valid link is provided
    if (!response.link || response.link.trim() === "") {
        alert("Error: The file could not be processed or the link is invalid.");
        return; // Exit if the response link is invalid
    }

    // Extract the link from the server response
    var responseLink = response.link;

    // Create a line break element
    var breakElement = document.createElement('br');

    // Create a new anchor element for the file download link
    var linkElement = document.createElement('a');
    linkElement.href = responseLink;                  // Set the hyperlink to the response link
    linkElement.textContent = "Download your file: " + responseLink; // Display text for the link

    // Get the container where links will be displayed
    var linkContainer = document.getElementById("sfs-links");
    
    // Append the new link and a line break to the container
    linkContainer.appendChild(linkElement);
    linkContainer.appendChild(breakElement);
});