import React, { useContext, useEffect, useState } from 'react';
import getArticles from './getArticles';
import Article from './interfaces/articles';
import uploadFile from './uploadFiles';
import { useNavigate } from 'react-router-dom';

interface BlogContextType {
    articles: Article[];
    routeChange: (title: string) => void;
}

const BlogContext = React.createContext<BlogContextType>({
    articles: [],
    routeChange: () => '',
});

export const useBlogContext = () => {
    return useContext(BlogContext);
};

const BlogContextProvider: React.FC<{ children: React.ReactNode }> = ({
    children,
}) => {
    const [articles, setArticles] = useState<Article[]>([]);

    let navigate = useNavigate();
    // Definir ici le upload

    // Load les articles au chargement du context
    // Mais il faut aussi les reload quand on ajoute un article
    useEffect(() => {
        (async () => {
            setArticles(await getArticles());
        })();
    }, [uploadFile]);
    // Not enough mais a verifier
    const routeChange = (title: string) => {
        let path = `/article/${title}`;
        navigate(path);
    };

    const contextValue = {
        articles,
        routeChange,
    };
    return (
        <BlogContext.Provider value={contextValue}>
            {children}
        </BlogContext.Provider>
    );
};

export default BlogContextProvider;
