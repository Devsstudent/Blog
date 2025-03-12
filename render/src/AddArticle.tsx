import { useState } from 'react';
import { Button } from './components/ui/button';
import { Input } from './components/ui/input';
import UploadComponent from './UploadComponent';
import uploadFile from './uploadFiles';

export default function AddArticle() {
    const [markdownContent, setMarkdownContent] = useState('');
    const [title, setTitle] = useState('');

    const [file, setFile] = useState<File | null>(null);

    const setUploadedFile = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = (e.target.files ?? [])[0] ?? null;
        if (file) {
            setFile(file);
        }
    };

    return (
        <>
            <div className="flex justify-evenly">
                <div className="ml-4 p-2 border-1 rounded-md">
                    <p className="flex justify-center p-5 text-xl">
                        Writing articles
                    </p>
                    <label className="flex text-xl mb-2"> Title </label>
                    <Input
                        className="flex"
                        onChange={(e) => {
                            setTitle(e.target.value);
                        }}
                    />
                    <div className="flex mb-2 text-xl">
                        <label> Content :</label>
                    </div>
                    <textarea
                        className="flex border"
                        placeholder="Start to type markdown here"
                        cols={70}
                        rows={15}
                        onChange={(e) => {
                            setMarkdownContent(e.target.value);
                        }}
                    />
                    <div className="flex justify-end m-2">
                        <Button
                            className="bg-blue-400 hover:cursor-pointer hover:bg-blue-500 shadow-md rounded-xl"
                            onClick={async () => {
                                const ret = await fetch(
                                    'http://localhost:8080/api/addArticleText',
                                    {
                                        method: 'POST',
                                        headers: {
                                            Accept: '*/*',
                                            'Content-Type': 'application/json',
                                        },
                                        body: JSON.stringify({
                                            title: title,
                                            content: markdownContent,
                                        }),
                                    }
                                );
                            }}
                        >
                            {' '}
                            Post
                        </Button>
                    </div>
                    {
                        // Maybe add a preview on the right
                    }
                </div>
                <div className="border-1 rounded-md shadow-md">
                    <p className="flex justify-center text-xl p-2">
                        Upload a file
                    </p>
                    <div className="flex w-auto items-center p-2">
                        <UploadComponent onChange={setUploadedFile} />
                        <span className="m-1">
                            Selected filed: {file?.name ?? 'None'}
                        </span>
                        <Button
                            className="hover:cursor-pointer shadow-md bg-emerald-200 hover:bg-emerald-300"
                            onClick={() => uploadFile(file)}
                        >
                            Upload
                        </Button>
                    </div>
                </div>
            </div>
        </>
    );
}
