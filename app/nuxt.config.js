import colors from 'vuetify/es5/util/colors'

export default {
    head: {
        title: 'Sports Kernel',
        meta: [
            {charset: 'utf-8'},
            {name: 'viewport', content: 'width=device-width, initial-scale=1'},
            {hid: 'description', name: 'description', content: ''},
            {name: 'format-detection', content: 'telephone=no'}
        ],
        link: [
            {rel: 'icon', type: 'image/x-icon', href: '/svg/color_logo-transparent.svg'}
        ]
    },

    css: [],

    plugins: [],

    components: true,

    buildModules: [
        '@nuxtjs/vuetify',
        '@nuxtjs/dotenv'
    ],

    modules: [
        '@nuxtjs/axios',
        '@nuxtjs/pwa',
        '@nuxtjs/auth-next',
        '@nuxtjs/apollo'
    ],

    auth: {
        strategies: {
            auth0: {
                domain: 'sports-kernel.us.auth0.com',
                clientId: 'YFF0nlkWCMbSDzWHiUxd7ZnHKlp9NUmD',
                audience: 'https://sports-kernel.us.auth0.com/api/v2/'
            }
        }
    },

    // Axios module configuration: https://go.nuxtjs.dev/config-axios
    axios: {
        common: {
            'Access-Control-Allow-Origin': '*'
        }
        // Access-Control-Allow-Origin: *
    },

    apollo: {
        tokenName: "auth._token.auth0", // specify token name
        // defaultOptions: {
        //     $query: {
        //         fetchPolicy: "network-only",
        //         errorPolicy: "all"
        //     }
        // },

        // watchLoading: "@/apollo/loadingHandler.js",
        // errorHandler: "@/apollo/errorHandler.js",
        clientConfigs: {
            default: {
                // httpEndpoint: process.env.GRAPHQL_BASE
                httpEndpoint: 'http://localhost:8080/graphql',
                httpLinkOptions: {
                    credentials: 'same-origin'
                },
            }
        }
    },


    // PWA module configuration: https://go.nuxtjs.dev/pwa
    pwa: {
        manifest: {
            lang: 'en'
        }
    },

    // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
    vuetify: {
        customVariables: ['~/assets/variables.scss'],
        theme: {
            dark: true,
            themes: {
                dark: {
                    primary: colors.blue.darken2,
                    accent: colors.grey.darken3,
                    secondary: colors.amber.darken3,
                    info: colors.teal.lighten1,
                    warning: colors.amber.base,
                    error: colors.deepOrange.accent4,
                    success: colors.green.accent3
                }
            }
        }
    },

    // Build Configuration: https://go.nuxtjs.dev/config-build
    build: {}
}
