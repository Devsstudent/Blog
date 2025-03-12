import Article from './interfaces/articles';

const PreviewArticle = ({ article }: { article: Article }) => {
    return (
        <>
            <div className="border-1 bg-gray-200/80 flex w-[80%] rounded-xl shadow-xl p-3 ">
                <p>
                    <strong>{article.title}</strong>
                    <br />{' '}
                    {article.content.length > 80
                        ? article.content.slice(0, 80) + '...'
                        : article.content}
                </p>
            </div>
        </>
    );
};

export default PreviewArticle;
