function removeExtension(filename: string) {
    return filename.substring(0, filename.lastIndexOf('.')) || filename;
}

const uploadFile = async (file: File | null) => {
    if (file) {
        const formData = new FormData();
        formData.append('file', file);
        formData.append('title', removeExtension(file.name));
        await fetch('http://localhost:8080/api/uploadArticle', {
            method: 'POST',
            body: formData,
        });
    }
};

export default uploadFile;
