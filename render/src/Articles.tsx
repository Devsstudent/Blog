import { useNavigate } from 'react-router-dom';
import { useBlogContext } from './BlogContext';
import PreviewArticle from './PreviewArticle';

const Articles = () => {
    const { articles } = useBlogContext();

    let navigate = useNavigate();
    const routeChange = (title: string) => {
        let path = `/article/${title}`;
        navigate(path);
    };

    return (
        <>
            <div className="p-5 m-5">
                <div className="text-2xl font-bold mb-3 flex justify-center">
                    Liste des articles
                </div>
                {articles?.map((article) => {
                    return (
                        <div
                            className="mb-5 cursor-pointer flex justify-center"
                            onClick={() => {
                                console.log(article.title);
                                return routeChange(article.title);
                            }}
                        >
                            <PreviewArticle article={article} />
                        </div>
                    );
                })}
            </div>
        </>
    );
};
export default Articles;
