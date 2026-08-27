import { postUploads } from "../api";
import type { CreateUploadRequest, UploadUrls } from "../api";
import type { ApiResult } from "@/helpers/apiResult";

export class UploadsService {
    /**
     * POST /uploads
     * Requests a signed URL to upload an image directly to storage, plus
     * the public URL it will be reachable at once the upload completes.
     */
    public static async createUpload(body: CreateUploadRequest): Promise<ApiResult<UploadUrls>> {
        const { data, error, response } = await postUploads({ body });

        return { data, error, status: response.status } as ApiResult<UploadUrls>;
    }
}
