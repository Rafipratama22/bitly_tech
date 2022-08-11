// import axios from "../../url/axios";
import axios from "axios"

export const ActionTypeUser = {
    isLogin: "IS_LOGIN_USER",
    isRegis: "IS_REGIS_USER",
    isUrl: "IS_URL",
    allProducts: "ALL_PRODUCTS",
    updateProduct: "UPDATE_PRODUCT",
    createProduct: "CREATE_PRODUCT"
};

// export const allChat = (payload) => {
//     return {
//         type: ActionTypeChat.allChat,
//         payload
//     }
// }

// export const filter = (payload) => {
//     return {
//         type: ActionTypeUserDoctor.filterUserDoctor,
//         payload
//     }
// }

export const isLogin = (payload) => {
    console.log(payload)
    return {
    type: ActionTypeUser.isLogin,
    payload
    }
}

export const isRegis = (payload) => {
    return {
        type: ActionTypeUser.isRegis,
        payload
    }
}

export const isUrl = (payload) => {
    return {
        type: ActionTypeUser.isUrl,
        payload
    }
}

export const allProducts = (payload) => {
    return {
        type: ActionTypeUser.allProducts,
        payload
    }
}

export const updateProduct = (payload) => {
    return {
        type: ActionTypeUser.updateProduct,
        payload
    }
}

export const createProduct = (payload) => {
    return {
        type: ActionTypeUser.createProduct,
        payload
    }
}

export const loginUser = (user) => async (dispatch) => {
    try {
        const {data} = await axios.post('http://localhost:8081/api/bitly/login', user)
        const access_token = data.access_token
        localStorage.setItem('access_token', access_token )
        dispatch(isLogin(true))
    } catch (err) {
        console.log(err)
    }
}


export const regisUser = (user) => async (dispatch) => {
    try {
        const {data} = await axios.post('http://localhost:8081/api/bitly/register', user,{
            headers: {'Content-Type': 'application/json'}
        })
        console.log(data)
        dispatch(isRegis(true))
    } catch (err) {
        console.log(err)
    }
}

export const makeUrl = (payload) => async (dispatch) => {
    try {
        const {data} = await axios.post('http://localhost:8081/api/bitly/create', payload,{
            headers: {'Content-Type': 'application/json'}
        })
        console.log(data)
        dispatch(isUrl(data.code))
    }catch (err) {
        console.log(err, "errr")
    }
}

export const fetchProducts = (payload) => async (dispatch) => {
    try {
        const access_token = localStorage.getItem("access_token")
        const {data} = await axios({
            method: "GET",
            url: "http://localhost:8081/api/bitly/products",
            headers: {
                Authorization: ""
            }
        })
        console.log(data)
        dispatch(allProducts(data))
    }catch(err) {
        console.log(err, "errr")
    }
}

export const updateRecentProduct = (payload) => async (dispatch) => {
    try {
        const {data} = await axios.put('http://localhost:8081/api/bitly/link/' + payload.id, payload, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': access_token
        }})
        console.log(data)
        dispatch(updateProduct(data))
    }catch(err) {
        console.log(err)
    }
}

export const createProductUser = (payload) => async (dispatch) => {
    try {
        const {data} = await axios.put('http://localhost:8081/api/bitly/link/', payload, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': access_token
        }})
        console.log(data)
        dispatch(createProduct(data))
    }catch (err) {
        console.log(err)
    }
}