function productExceptSelf(nums: number[]): number[] {
    const n: Number = nums.length;
    const result: number[] = new Array(n);
    
    // Calculate left products
    result[0] = 1;
    for (let i: number = 1; i < n; i++) {
        result[i] = result[i - 1] * nums[i - 1];
    }
    
    // Calculate right products and multiply with left products
    let rightProduct: number = 1;
    for (let i: number = n - 1; i >= 0; i--) {
        result[i] = result[i] * rightProduct;
        rightProduct *= nums[i];
    }
    
    return result;
}