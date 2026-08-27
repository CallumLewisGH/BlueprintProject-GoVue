import { 
    getUsersMe, 
    postUsersSearch, 
    patchUsersMe, 
    deleteUsersMe,
    getUsersById,
} from "../api";
import type { 
    UserPrivateProfile, 
    UserPublicProfile,
    UpdateUserRequest, 
    QueryParameter
} from "../api";
import { QueryBuilder } from "../helpers/queryBuilder";
import type { ApiResult } from "@/helpers/apiResult";

export class UsersService {
    /**
     * POST /users/search
     * Fetches a list of public profiles based on query filters.
     */
    public static async getUsers(query?: QueryParameter[]): Promise<ApiResult<UserPublicProfile[]>> {
        if (!query) {
            query = new QueryBuilder().build();
        }

        const { data, error, response } = await postUsersSearch({ body: query });

        return { data, error, status: response.status } as ApiResult<UserPublicProfile[]>;
    }

    /**
     * GET /users/me
     * Fetches the private profile of the currently authenticated user.
     */
    public static async getCurrentUser(): Promise<ApiResult<UserPrivateProfile>> {
        const { data, error, response } = await getUsersMe();

        return { data, error, status: response.status } as ApiResult<UserPrivateProfile>;
    }

    /**
     * GET /users/{id}
     * Fetches a public profile by its specific GUID.
     */
    public static async getUserById(id: string): Promise<ApiResult<UserPublicProfile>> {
        const { data, error, response } = await getUsersById({ path: { id } });

        return { data, error, status: response.status } as ApiResult<UserPublicProfile>;
    }

    /**
     * PATCH /users/me
     * Updates the authenticated user's profile.
     */
    public static async updateCurrentUser(body: UpdateUserRequest): Promise<ApiResult<UserPrivateProfile>> {
        const { data, error, response } = await patchUsersMe({ body: body });

        return { data, error, status: response.status } as ApiResult<UserPrivateProfile>;
    }

    /**
     * DELETE /users/me
     * Deletes the authenticated user's account.
     */
    public static async deleteCurrentUser(): Promise<ApiResult<UserPrivateProfile>> {
        const { data, error, response } = await deleteUsersMe();

        return { data, error, status: response.status } as ApiResult<UserPrivateProfile>;
    }
}