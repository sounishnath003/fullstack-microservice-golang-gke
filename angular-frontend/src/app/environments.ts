// Read here: https://dev.to/dimeloper/managing-environment-variables-in-angular-apps-14gn

console.log({ AUTH_SERVICE_ENDPOINT: import.meta.env });

export const NG_APP_AUTH_SERVICE_ENDPOINT = import.meta.env.NG_APP_AUTH_SERVICE_ENDPOINT;
export const NG_APP_BLOGS_SERVICE_ENDPOINT = import.meta.env.NG_APP_BLOGS_SERVICE_ENDPOINT;