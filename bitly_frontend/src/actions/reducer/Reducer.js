import { ActionTypeUser } from "../../actions/actions/Action";

const initialState = {
  allProducts: [
    {
      ID: 1,
      DestinationUrl: "",
      ModifyUrl: "",
      Title: "",
      Click: 0,
      UserID: 4,
      Source: "",
      Medium: "",
      Campaign: "",
      CreatedAt: "",
      UpdatedAt: ""
    },
  ],
  isLogin: "",
  isRegis: "",
  isUrl: "",
  updateProduct: {
      ID: 1,
      DestinationUrl: "",
      ModifyUrl: "",
      Title: "",
      Click: 0,
      UserID: 3,
      Source: "",
      Medium: "",
      Campaign: "",
      CreatedAt: "",
      UpdatedAt: ""
  },
  createProduct: {
      ID: 3,
      DestinationUrl: "https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/",
      ModifyUrl: "baiCMR",
      Title: "",
      Click: 0,
      UserID: 3,
      Source: "",
      Medium: "",
      Campaign: "",
      CreatedAt: "2022-08-10T09:45:37.2592156+07:00",
      UpdatedAt: "2022-08-10T09:45:37.2592156+07:00"
  }
};

export default function ReducerUser(state = initialState, action) {
  switch (action.type) {
    case ActionTypeUser.allProducts:
      return {...state, allProducts: action.payload};
    case ActionTypeUser.isLogin:
      console.log(action.payload);
      return { ...state, isLogin: action.payload };
    case ActionTypeUser.isRegis:
      return { ...state, isRegis: action.payload }
    case ActionTypeUser.isUrl:
      return {...state, isUrl: action.payload}
    case ActionTypeUser.updateProduct:
      return {...state, updateProduct: action.payload}
    case ActionTypeUser.createProduct:
      return {...state, createProduct: action.payload}
    default:
      return state;
  }
}
