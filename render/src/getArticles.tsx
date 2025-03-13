import Article from './interfaces/articles';

export default async function getArticles() {
    const ret = await fetch('http://localhost:8080/api/getAllArticles', {
        method: 'GET',
        headers: {
            Accept: '*/*',
            'Content-Type': 'application/json',
        },
    });
    const articles: Article[] = await ret.json();
    if (articles.length > 0) {
        return articles.filter((article) => article.validated);
    }
    return [];
}
