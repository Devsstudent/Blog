import { useEffect, useState } from 'react';
import Article from './interfaces/articles';

const RenderArticle = () => {
    //faudra que ce soit un id plus tard
    const url = window.location.href;
    const regex = /\/([^\/\?#]*)(?:[?#].*)?$/;
    const match = url.match(regex);
    const paramsTitle = match ? match[1] : null;

    const [article, setArticle] = useState<Article>();

    useEffect(() => {
        async function fetchArticles() {
            const ret = await fetch(
                `http://localhost:8080/api/getArticleFromTitle?title=${paramsTitle}`,
                {
                    method: 'GET',
                    headers: {
                        Accept: '*/*',
                        'Content-Type': 'application/json',
                    },
                }
            );
            const article: Article = await ret.json();
            setArticle(article);
        }
        fetchArticles();
    }, []);

    return (
        <>
            <div className="flex justify-center py-3">
                {article?.html && (
                    <div
                        className="prose"
                        dangerouslySetInnerHTML={{ __html: article?.html }}
                    />
                )}
            </div>
        </>
    );
};
export default RenderArticle;
