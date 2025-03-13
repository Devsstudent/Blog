import { useState } from 'react';
import { useBlogContext } from './BlogContext';
import PreviewArticle from './PreviewArticle';
import AddArticle from './AddArticle';

export default function Home() {
    const [file, setFile] = useState<File | null>(null);

    const setUploadedFile = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = (e.target.files ?? [])[0] ?? null;
        if (file) {
            setFile(file);
        }
    };
    const { articles, routeChange } = useBlogContext();

    return (
        <div className="mt-10">
            <div className="flex justify-center">
                <div className="text-xl font-serif shadow-md p-2 m-2">
                    <p>
                        Hello, this blog is just a portfolio.
                        <br />I put articles about my work and hobbie.
                        <br />
                        The style should improve with the time !
                    </p>
                </div>
            </div>
            {articles.length > 0 && (
                <span className="text-2xl ml-10 flex justify-center">
                    Recent articles
                </span>
            )}
            {articles
                ?.map((article) => {
                    return (
                        <div
                            className="mb-5 cursor-pointer flex justify-center"
                            onClick={() => {
                                return routeChange(article.title);
                            }}
                        >
                            <PreviewArticle article={article} />
                        </div>
                    );
                })
                .slice(0, 3)}

            <AddArticle />
        </div>
    );
}
