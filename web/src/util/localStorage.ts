
export const TOKENKEY = 'sample';
// set token to local storage
// returns void
export const setTokenStorage = (token: string) => {
    localStorage.setItem(TOKENKEY, token);
};
// get token from local storage
// returns string or undefined
// the function is not async because local storage is sync
export const getTokenStorage = (): string | undefined => {
    const token = localStorage.getItem(TOKENKEY);
    if (token) {
        return token;
    }
    return undefined;
}