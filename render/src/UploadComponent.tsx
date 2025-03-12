const UploadComponent = ({
    onChange,
}: {
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) => {
    return (
        <>
            <label
                className="p-3 cursor-pointer mx-4 my-4 bg-blue-100 border-1 border-solid rounded-2xl hover:bg-blue-200 "
                htmlFor="fileInput"
            >
                + Select a file
            </label>
            <input
                className="hidden"
                accept=".md,text/markdown"
                type="file"
                id="fileInput"
                onChange={onChange}
            ></input>
        </>
    );
};

export default UploadComponent;
