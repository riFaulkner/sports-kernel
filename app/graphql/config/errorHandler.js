export default (graphqlError, { store, error, redirect, route }) => {
    console.log({ graphqlError })

    const { networkError, message, gqlError, graphqlErrors } = graphqlError

    // handle error
    console.log("what should I do with this error?");
    if (message === "GraphQL error: Access denied") {
        console.log("attempting to dispatch message to global toast")
        store.dispatch("application/alertError", {message: "Unable to load, required permissions not found"});
    }
}
