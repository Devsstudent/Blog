import { useState } from "react";
import { Button } from "./components/ui/button";

export default function AddArticle() {
    const [markdownContent, setMarkdownContent] = useState("");

    return (
        <>
            <textarea cols={70} rows={30} onChange={(e) => {
                setMarkdownContent(e.target.value);
            }} />
            <Button className="bg-blue-400" onClick={async () => {
                const ret = await fetch('http://localhost:8080/api/addArticleText', {
                    method: 'POST', headers: {
                        'Accept': '*/*',
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        'title': "brain",
                        'content': markdownContent
                    })
                })
                console.log('sent to the back', ret.formData);
            }}> Click MEE</Button >
        </>
    )
}
