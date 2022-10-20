
#if !defined(WASMER_H_MACROS)

#define WASMER_H_MACROS

// Define the `ARCH_X86_X64` constant.
#if defined(MSVC) && defined(_M_AMD64)
#  define ARCH_X86_64
#elif (defined(GCC) || defined(__GNUC__) || defined(__clang__)) && defined(__x86_64__)
#  define ARCH_X86_64
#endif

// Compatibility with non-Clang compilers.
#if !defined(__has_attribute)
#  define __has_attribute(x) 0
#endif

// Compatibility with non-Clang compilers.
#if !defined(__has_declspec_attribute)
#  define __has_declspec_attribute(x) 0
#endif

// Define the `DEPRECATED` macro.
#if defined(GCC) || defined(__GNUC__) || __has_attribute(deprecated)
#  define DEPRECATED(message) __attribute__((deprecated(message)))
#elif defined(MSVC) || __has_declspec_attribute(deprecated)
#  define DEPRECATED(message) __declspec(deprecated(message))
#endif

#endif // WASMER_H_MACROS


#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

/**
 * The `wasmer_result_t` enum is a type that represents either a
 * success, or a failure.
 */
typedef enum {
  /**
   * Represents a success.
   */
  VM_EXEC_OK = 1,
  /**
   * Represents a failure.
   */
  VM_EXEC_ERROR = 2,
} vm_exec_result_t;

/**
 * Opaque pointer to a `wasmer_runtime::Instance` value in Rust.
 *
 * A `wasmer_runtime::Instance` represents a WebAssembly instance. It
 * is generally generated by the `wasmer_instantiate()` function, or by
 * the `wasmer_module_instantiate()` function for the most common paths.
 */
typedef struct {

} vm_exec_instance_t;

typedef struct {

} vm_exec_executor_t;

typedef struct {
  int64_t (*get_gas_left_func_ptr)(void *context);
  void (*get_sc_address_func_ptr)(void *context, int32_t result_offset);
  void (*get_owner_address_func_ptr)(void *context, int32_t result_offset);
  int32_t (*get_shard_of_address_func_ptr)(void *context, int32_t address_offset);
  int32_t (*is_smart_contract_func_ptr)(void *context, int32_t address_offset);
  void (*signal_error_func_ptr)(void *context, int32_t message_offset, int32_t message_length);
  void (*get_external_balance_func_ptr)(void *context, int32_t address_offset, int32_t result_offset);
  int32_t (*get_block_hash_func_ptr)(void *context, int64_t nonce, int32_t result_offset);
  int32_t (*get_esdt_balance_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce, int32_t result_offset);
  int32_t (*get_esdt_nft_name_length_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce);
  int32_t (*get_esdt_nft_attribute_length_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce);
  int32_t (*get_esdt_nft_uri_length_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce);
  int32_t (*get_esdt_token_data_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce, int32_t value_handle, int32_t properties_offset, int32_t hash_offset, int32_t name_offset, int32_t attributes_offset, int32_t creator_offset, int32_t royalties_handle, int32_t uris_offset);
  int64_t (*get_esdt_local_roles_func_ptr)(void *context, int32_t token_id_handle);
  int32_t (*validate_token_identifier_func_ptr)(void *context, int32_t token_id_handle);
  int32_t (*transfer_value_func_ptr)(void *context, int32_t dest_offset, int32_t value_offset, int32_t data_offset, int32_t length);
  int32_t (*transfer_value_execute_func_ptr)(void *context, int32_t dest_offset, int32_t value_offset, int64_t gas_limit, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*transfer_esdt_execute_func_ptr)(void *context, int32_t dest_offset, int32_t token_id_offset, int32_t token_id_len, int32_t value_offset, int64_t gas_limit, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*transfer_esdt_nft_execute_func_ptr)(void *context, int32_t dest_offset, int32_t token_id_offset, int32_t token_id_len, int32_t value_offset, int64_t nonce, int64_t gas_limit, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*multi_transfer_esdt_nft_execute_func_ptr)(void *context, int32_t dest_offset, int32_t num_token_transfers, int32_t token_transfers_args_length_offset, int32_t token_transfer_data_offset, int64_t gas_limit, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*create_async_call_func_ptr)(void *context, int32_t dest_offset, int32_t value_offset, int32_t data_offset, int32_t data_length, int32_t success_offset, int32_t success_length, int32_t error_offset, int32_t error_length, int64_t gas, int64_t extra_gas_for_callback);
  int32_t (*set_async_context_callback_func_ptr)(void *context, int32_t callback, int32_t callback_length, int32_t data, int32_t data_length, int64_t gas);
  void (*upgrade_contract_func_ptr)(void *context, int32_t dest_offset, int64_t gas_limit, int32_t value_offset, int32_t code_offset, int32_t code_metadata_offset, int32_t length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  void (*upgrade_from_source_contract_func_ptr)(void *context, int32_t dest_offset, int64_t gas_limit, int32_t value_offset, int32_t source_contract_address_offset, int32_t code_metadata_offset, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  void (*delete_contract_func_ptr)(void *context, int32_t dest_offset, int64_t gas_limit, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  void (*async_call_func_ptr)(void *context, int32_t dest_offset, int32_t value_offset, int32_t data_offset, int32_t length);
  int32_t (*get_argument_length_func_ptr)(void *context, int32_t id);
  int32_t (*get_argument_func_ptr)(void *context, int32_t id, int32_t arg_offset);
  int32_t (*get_function_func_ptr)(void *context, int32_t function_offset);
  int32_t (*get_num_arguments_func_ptr)(void *context);
  int32_t (*storage_store_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t data_offset, int32_t data_length);
  int32_t (*storage_load_length_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int32_t (*storage_load_from_address_func_ptr)(void *context, int32_t address_offset, int32_t key_offset, int32_t key_length, int32_t data_offset);
  int32_t (*storage_load_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t data_offset);
  int32_t (*set_storage_lock_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int64_t lock_timestamp);
  int64_t (*get_storage_lock_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int32_t (*is_storage_locked_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int32_t (*clear_storage_lock_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  void (*get_caller_func_ptr)(void *context, int32_t result_offset);
  void (*check_no_payment_func_ptr)(void *context);
  int32_t (*get_call_value_func_ptr)(void *context, int32_t result_offset);
  int32_t (*get_esdt_value_func_ptr)(void *context, int32_t result_offset);
  int32_t (*get_esdt_value_by_index_func_ptr)(void *context, int32_t result_offset, int32_t index);
  int32_t (*get_esdt_token_name_func_ptr)(void *context, int32_t result_offset);
  int32_t (*get_esdt_token_name_by_index_func_ptr)(void *context, int32_t result_offset, int32_t index);
  int64_t (*get_esdt_token_nonce_func_ptr)(void *context);
  int64_t (*get_esdt_token_nonce_by_index_func_ptr)(void *context, int32_t index);
  int64_t (*get_current_esdt_nft_nonce_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len);
  int32_t (*get_esdt_token_type_func_ptr)(void *context);
  int32_t (*get_esdt_token_type_by_index_func_ptr)(void *context, int32_t index);
  int32_t (*get_num_esdt_transfers_func_ptr)(void *context);
  int32_t (*get_call_value_token_name_func_ptr)(void *context, int32_t call_value_offset, int32_t token_name_offset);
  int32_t (*get_call_value_token_name_by_index_func_ptr)(void *context, int32_t call_value_offset, int32_t token_name_offset, int32_t index);
  void (*write_log_func_ptr)(void *context, int32_t data_pointer, int32_t data_length, int32_t topic_ptr, int32_t num_topics);
  void (*write_event_log_func_ptr)(void *context, int32_t num_topics, int32_t topic_lengths_offset, int32_t topic_offset, int32_t data_offset, int32_t data_length);
  int64_t (*get_block_timestamp_func_ptr)(void *context);
  int64_t (*get_block_nonce_func_ptr)(void *context);
  int64_t (*get_block_round_func_ptr)(void *context);
  int64_t (*get_block_epoch_func_ptr)(void *context);
  void (*get_block_random_seed_func_ptr)(void *context, int32_t pointer);
  void (*get_state_root_hash_func_ptr)(void *context, int32_t pointer);
  int64_t (*get_prev_block_timestamp_func_ptr)(void *context);
  int64_t (*get_prev_block_nonce_func_ptr)(void *context);
  int64_t (*get_prev_block_round_func_ptr)(void *context);
  int64_t (*get_prev_block_epoch_func_ptr)(void *context);
  void (*get_prev_block_random_seed_func_ptr)(void *context, int32_t pointer);
  void (*finish_func_ptr)(void *context, int32_t pointer, int32_t length);
  int32_t (*execute_on_same_context_func_ptr)(void *context, int64_t gas_limit, int32_t address_offset, int32_t value_offset, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*execute_on_dest_context_func_ptr)(void *context, int64_t gas_limit, int32_t address_offset, int32_t value_offset, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*execute_read_only_func_ptr)(void *context, int64_t gas_limit, int32_t address_offset, int32_t function_offset, int32_t function_length, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*create_contract_func_ptr)(void *context, int64_t gas_limit, int32_t value_offset, int32_t code_offset, int32_t code_metadata_offset, int32_t length, int32_t result_offset, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*deploy_from_source_contract_func_ptr)(void *context, int64_t gas_limit, int32_t value_offset, int32_t source_contract_address_offset, int32_t code_metadata_offset, int32_t result_address_offset, int32_t num_arguments, int32_t arguments_length_offset, int32_t data_offset);
  int32_t (*get_num_return_data_func_ptr)(void *context);
  int32_t (*get_return_data_size_func_ptr)(void *context, int32_t result_id);
  int32_t (*get_return_data_func_ptr)(void *context, int32_t result_id, int32_t data_offset);
  void (*clean_return_data_func_ptr)(void *context);
  void (*delete_from_return_data_func_ptr)(void *context, int32_t result_id);
  void (*get_original_tx_hash_func_ptr)(void *context, int32_t data_offset);
  void (*get_current_tx_hash_func_ptr)(void *context, int32_t data_offset);
  void (*get_prev_tx_hash_func_ptr)(void *context, int32_t data_offset);
  void (*managed_sc_address_func_ptr)(void *context, int32_t destination_handle);
  void (*managed_owner_address_func_ptr)(void *context, int32_t destination_handle);
  void (*managed_caller_func_ptr)(void *context, int32_t destination_handle);
  void (*managed_signal_error_func_ptr)(void *context, int32_t err_handle);
  void (*managed_write_log_func_ptr)(void *context, int32_t topics_handle, int32_t data_handle);
  void (*managed_get_original_tx_hash_func_ptr)(void *context, int32_t result_handle);
  void (*managed_get_state_root_hash_func_ptr)(void *context, int32_t result_handle);
  void (*managed_get_block_random_seed_func_ptr)(void *context, int32_t result_handle);
  void (*managed_get_prev_block_random_seed_func_ptr)(void *context, int32_t result_handle);
  void (*managed_get_return_data_func_ptr)(void *context, int32_t result_id, int32_t result_handle);
  void (*managed_get_multi_esdt_call_value_func_ptr)(void *context, int32_t multi_call_value_handle);
  void (*managed_get_esdt_balance_func_ptr)(void *context, int32_t address_handle, int32_t token_id_handle, int64_t nonce, int32_t value_handle);
  void (*managed_get_esdt_token_data_func_ptr)(void *context, int32_t address_handle, int32_t token_id_handle, int64_t nonce, int32_t value_handle, int32_t properties_handle, int32_t hash_handle, int32_t name_handle, int32_t attributes_handle, int32_t creator_handle, int32_t royalties_handle, int32_t uris_handle);
  void (*managed_async_call_func_ptr)(void *context, int32_t dest_handle, int32_t value_handle, int32_t function_handle, int32_t arguments_handle);
  int32_t (*managed_create_async_call_func_ptr)(void *context, int32_t dest_handle, int32_t value_handle, int32_t function_handle, int32_t arguments_handle, int32_t success_offset, int32_t success_length, int32_t error_offset, int32_t error_length, int64_t gas, int64_t extra_gas_for_callback, int32_t callback_closure_handle);
  void (*managed_get_callback_closure_func_ptr)(void *context, int32_t callback_closure_handle);
  void (*managed_upgrade_from_source_contract_func_ptr)(void *context, int32_t dest_handle, int64_t gas, int32_t value_handle, int32_t address_handle, int32_t code_metadata_handle, int32_t arguments_handle, int32_t result_handle);
  void (*managed_upgrade_contract_func_ptr)(void *context, int32_t dest_handle, int64_t gas, int32_t value_handle, int32_t code_handle, int32_t code_metadata_handle, int32_t arguments_handle, int32_t result_handle);
  void (*managed_delete_contract_func_ptr)(void *context, int32_t dest_handle, int64_t gas_limit, int32_t arguments_handle);
  int32_t (*managed_deploy_from_source_contract_func_ptr)(void *context, int64_t gas, int32_t value_handle, int32_t address_handle, int32_t code_metadata_handle, int32_t arguments_handle, int32_t result_address_handle, int32_t result_handle);
  int32_t (*managed_create_contract_func_ptr)(void *context, int64_t gas, int32_t value_handle, int32_t code_handle, int32_t code_metadata_handle, int32_t arguments_handle, int32_t result_address_handle, int32_t result_handle);
  int32_t (*managed_execute_read_only_func_ptr)(void *context, int64_t gas, int32_t address_handle, int32_t function_handle, int32_t arguments_handle, int32_t result_handle);
  int32_t (*managed_execute_on_same_context_func_ptr)(void *context, int64_t gas, int32_t address_handle, int32_t value_handle, int32_t function_handle, int32_t arguments_handle, int32_t result_handle);
  int32_t (*managed_execute_on_dest_context_func_ptr)(void *context, int64_t gas, int32_t address_handle, int32_t value_handle, int32_t function_handle, int32_t arguments_handle, int32_t result_handle);
  int32_t (*managed_multi_transfer_esdt_nft_execute_func_ptr)(void *context, int32_t dst_handle, int32_t token_transfers_handle, int64_t gas_limit, int32_t function_handle, int32_t arguments_handle);
  int32_t (*managed_transfer_value_execute_func_ptr)(void *context, int32_t dst_handle, int32_t value_handle, int64_t gas_limit, int32_t function_handle, int32_t arguments_handle);
  int32_t (*managed_is_esdt_frozen_func_ptr)(void *context, int32_t address_handle, int32_t token_id_handle, int64_t nonce);
  int32_t (*managed_is_esdt_limited_transfer_func_ptr)(void *context, int32_t token_id_handle);
  int32_t (*managed_is_esdt_paused_func_ptr)(void *context, int32_t token_id_handle);
  void (*managed_buffer_to_hex_func_ptr)(void *context, int32_t source_handle, int32_t dest_handle);
  int32_t (*big_float_new_from_parts_func_ptr)(void *context, int32_t integral_part, int32_t fractional_part, int32_t exponent);
  int32_t (*big_float_new_from_frac_func_ptr)(void *context, int64_t numerator, int64_t denominator);
  int32_t (*big_float_new_from_sci_func_ptr)(void *context, int64_t significand, int64_t exponent);
  void (*big_float_add_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_float_sub_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_float_mul_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_float_div_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_float_neg_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  void (*big_float_clone_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  int32_t (*big_float_cmp_func_ptr)(void *context, int32_t op1_handle, int32_t op2_handle);
  void (*big_float_abs_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  int32_t (*big_float_sign_func_ptr)(void *context, int32_t op_handle);
  void (*big_float_sqrt_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  void (*big_float_pow_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle, int32_t exponent);
  void (*big_float_floor_func_ptr)(void *context, int32_t dest_big_int_handle, int32_t op_handle);
  void (*big_float_ceil_func_ptr)(void *context, int32_t dest_big_int_handle, int32_t op_handle);
  void (*big_float_truncate_func_ptr)(void *context, int32_t dest_big_int_handle, int32_t op_handle);
  void (*big_float_set_int64_func_ptr)(void *context, int32_t destination_handle, int64_t value);
  int32_t (*big_float_is_int_func_ptr)(void *context, int32_t op_handle);
  void (*big_float_set_big_int_func_ptr)(void *context, int32_t destination_handle, int32_t big_int_handle);
  void (*big_float_get_const_pi_func_ptr)(void *context, int32_t destination_handle);
  void (*big_float_get_const_e_func_ptr)(void *context, int32_t destination_handle);
  void (*big_int_get_unsigned_argument_func_ptr)(void *context, int32_t id, int32_t destination_handle);
  void (*big_int_get_signed_argument_func_ptr)(void *context, int32_t id, int32_t destination_handle);
  int32_t (*big_int_storage_store_unsigned_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t source_handle);
  int32_t (*big_int_storage_load_unsigned_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t destination_handle);
  void (*big_int_get_call_value_func_ptr)(void *context, int32_t destination_handle);
  void (*big_int_get_esdt_call_value_func_ptr)(void *context, int32_t destination);
  void (*big_int_get_esdt_call_value_by_index_func_ptr)(void *context, int32_t destination_handle, int32_t index);
  void (*big_int_get_external_balance_func_ptr)(void *context, int32_t address_offset, int32_t result);
  void (*big_int_get_esdt_external_balance_func_ptr)(void *context, int32_t address_offset, int32_t token_id_offset, int32_t token_id_len, int64_t nonce, int32_t result_handle);
  int32_t (*big_int_new_func_ptr)(void *context, int64_t small_value);
  int32_t (*big_int_unsigned_byte_length_func_ptr)(void *context, int32_t reference_handle);
  int32_t (*big_int_signed_byte_length_func_ptr)(void *context, int32_t reference_handle);
  int32_t (*big_int_get_unsigned_bytes_func_ptr)(void *context, int32_t reference_handle, int32_t byte_offset);
  int32_t (*big_int_get_signed_bytes_func_ptr)(void *context, int32_t reference_handle, int32_t byte_offset);
  void (*big_int_set_unsigned_bytes_func_ptr)(void *context, int32_t destination_handle, int32_t byte_offset, int32_t byte_length);
  void (*big_int_set_signed_bytes_func_ptr)(void *context, int32_t destination_handle, int32_t byte_offset, int32_t byte_length);
  int32_t (*big_int_is_int64_func_ptr)(void *context, int32_t destination_handle);
  int64_t (*big_int_get_int64_func_ptr)(void *context, int32_t destination_handle);
  void (*big_int_set_int64_func_ptr)(void *context, int32_t destination_handle, int64_t value);
  void (*big_int_add_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_sub_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_mul_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_tdiv_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_tmod_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_ediv_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_emod_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_sqrt_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  void (*big_int_pow_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  int32_t (*big_int_log2_func_ptr)(void *context, int32_t op1_handle);
  void (*big_int_abs_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  void (*big_int_neg_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  int32_t (*big_int_sign_func_ptr)(void *context, int32_t op_handle);
  int32_t (*big_int_cmp_func_ptr)(void *context, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_not_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle);
  void (*big_int_and_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_or_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_xor_func_ptr)(void *context, int32_t destination_handle, int32_t op1_handle, int32_t op2_handle);
  void (*big_int_shr_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle, int32_t bits);
  void (*big_int_shl_func_ptr)(void *context, int32_t destination_handle, int32_t op_handle, int32_t bits);
  void (*big_int_finish_unsigned_func_ptr)(void *context, int32_t reference_handle);
  void (*big_int_finish_signed_func_ptr)(void *context, int32_t reference_handle);
  void (*big_int_to_string_func_ptr)(void *context, int32_t big_int_handle, int32_t destination_handle);
  int32_t (*mbuffer_new_func_ptr)(void *context);
  int32_t (*mbuffer_new_from_bytes_func_ptr)(void *context, int32_t data_offset, int32_t data_length);
  int32_t (*mbuffer_get_length_func_ptr)(void *context, int32_t m_buffer_handle);
  int32_t (*mbuffer_get_bytes_func_ptr)(void *context, int32_t m_buffer_handle, int32_t result_offset);
  int32_t (*mbuffer_get_byte_slice_func_ptr)(void *context, int32_t source_handle, int32_t starting_position, int32_t slice_length, int32_t result_offset);
  int32_t (*mbuffer_copy_byte_slice_func_ptr)(void *context, int32_t source_handle, int32_t starting_position, int32_t slice_length, int32_t destination_handle);
  int32_t (*mbuffer_eq_func_ptr)(void *context, int32_t m_buffer_handle1, int32_t m_buffer_handle2);
  int32_t (*mbuffer_set_bytes_func_ptr)(void *context, int32_t m_buffer_handle, int32_t data_offset, int32_t data_length);
  int32_t (*mbuffer_set_byte_slice_func_ptr)(void *context, int32_t m_buffer_handle, int32_t starting_position, int32_t data_length, int32_t data_offset);
  int32_t (*mbuffer_append_func_ptr)(void *context, int32_t accumulator_handle, int32_t data_handle);
  int32_t (*mbuffer_append_bytes_func_ptr)(void *context, int32_t accumulator_handle, int32_t data_offset, int32_t data_length);
  int32_t (*mbuffer_to_big_int_unsigned_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_int_handle);
  int32_t (*mbuffer_to_big_int_signed_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_int_handle);
  int32_t (*mbuffer_from_big_int_unsigned_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_int_handle);
  int32_t (*mbuffer_from_big_int_signed_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_int_handle);
  int32_t (*mbuffer_to_big_float_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_float_handle);
  int32_t (*mbuffer_from_big_float_func_ptr)(void *context, int32_t m_buffer_handle, int32_t big_float_handle);
  int32_t (*mbuffer_storage_store_func_ptr)(void *context, int32_t key_handle, int32_t source_handle);
  int32_t (*mbuffer_storage_load_func_ptr)(void *context, int32_t key_handle, int32_t destination_handle);
  void (*mbuffer_storage_load_from_address_func_ptr)(void *context, int32_t address_handle, int32_t key_handle, int32_t destination_handle);
  int32_t (*mbuffer_get_argument_func_ptr)(void *context, int32_t id, int32_t destination_handle);
  int32_t (*mbuffer_finish_func_ptr)(void *context, int32_t source_handle);
  int32_t (*mbuffer_set_random_func_ptr)(void *context, int32_t destination_handle, int32_t length);
  int64_t (*small_int_get_unsigned_argument_func_ptr)(void *context, int32_t id);
  int64_t (*small_int_get_signed_argument_func_ptr)(void *context, int32_t id);
  void (*small_int_finish_unsigned_func_ptr)(void *context, int64_t value);
  void (*small_int_finish_signed_func_ptr)(void *context, int64_t value);
  int32_t (*small_int_storage_store_unsigned_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int64_t value);
  int32_t (*small_int_storage_store_signed_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int64_t value);
  int64_t (*small_int_storage_load_unsigned_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int64_t (*small_int_storage_load_signed_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int64_t (*int64get_argument_func_ptr)(void *context, int32_t id);
  void (*int64finish_func_ptr)(void *context, int64_t value);
  int32_t (*int64storage_store_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int64_t value);
  int64_t (*int64storage_load_func_ptr)(void *context, int32_t key_offset, int32_t key_length);
  int32_t (*sha256_func_ptr)(void *context, int32_t data_offset, int32_t length, int32_t result_offset);
  int32_t (*managed_sha256_func_ptr)(void *context, int32_t input_handle, int32_t output_handle);
  int32_t (*keccak256_func_ptr)(void *context, int32_t data_offset, int32_t length, int32_t result_offset);
  int32_t (*managed_keccak256_func_ptr)(void *context, int32_t input_handle, int32_t output_handle);
  int32_t (*ripemd160_func_ptr)(void *context, int32_t data_offset, int32_t length, int32_t result_offset);
  int32_t (*managed_ripemd160_func_ptr)(void *context, int32_t input_handle, int32_t output_handle);
  int32_t (*verify_bls_func_ptr)(void *context, int32_t key_offset, int32_t message_offset, int32_t message_length, int32_t sig_offset);
  int32_t (*managed_verify_bls_func_ptr)(void *context, int32_t key_handle, int32_t message_handle, int32_t sig_handle);
  int32_t (*verify_ed25519_func_ptr)(void *context, int32_t key_offset, int32_t message_offset, int32_t message_length, int32_t sig_offset);
  int32_t (*managed_verify_ed25519_func_ptr)(void *context, int32_t key_handle, int32_t message_handle, int32_t sig_handle);
  int32_t (*verify_custom_secp256k1_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t message_offset, int32_t message_length, int32_t sig_offset, int32_t hash_type);
  int32_t (*managed_verify_custom_secp256k1_func_ptr)(void *context, int32_t key_handle, int32_t message_handle, int32_t sig_handle, int32_t hash_type);
  int32_t (*verify_secp256k1_func_ptr)(void *context, int32_t key_offset, int32_t key_length, int32_t message_offset, int32_t message_length, int32_t sig_offset);
  int32_t (*managed_verify_secp256k1_func_ptr)(void *context, int32_t key_handle, int32_t message_handle, int32_t sig_handle);
  int32_t (*encode_secp256k1_der_signature_func_ptr)(void *context, int32_t r_offset, int32_t r_length, int32_t s_offset, int32_t s_length, int32_t sig_offset);
  int32_t (*managed_encode_secp256k1_der_signature_func_ptr)(void *context, int32_t r_handle, int32_t s_handle, int32_t sig_handle);
  void (*add_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t fst_point_xhandle, int32_t fst_point_yhandle, int32_t snd_point_xhandle, int32_t snd_point_yhandle);
  void (*double_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t point_xhandle, int32_t point_yhandle);
  int32_t (*is_on_curve_ec_func_ptr)(void *context, int32_t ec_handle, int32_t point_xhandle, int32_t point_yhandle);
  int32_t (*scalar_base_mult_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_offset, int32_t length);
  int32_t (*managed_scalar_base_mult_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_handle);
  int32_t (*scalar_mult_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t point_xhandle, int32_t point_yhandle, int32_t data_offset, int32_t length);
  int32_t (*managed_scalar_mult_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t point_xhandle, int32_t point_yhandle, int32_t data_handle);
  int32_t (*marshal_ec_func_ptr)(void *context, int32_t x_pair_handle, int32_t y_pair_handle, int32_t ec_handle, int32_t result_offset);
  int32_t (*managed_marshal_ec_func_ptr)(void *context, int32_t x_pair_handle, int32_t y_pair_handle, int32_t ec_handle, int32_t result_handle);
  int32_t (*marshal_compressed_ec_func_ptr)(void *context, int32_t x_pair_handle, int32_t y_pair_handle, int32_t ec_handle, int32_t result_offset);
  int32_t (*managed_marshal_compressed_ec_func_ptr)(void *context, int32_t x_pair_handle, int32_t y_pair_handle, int32_t ec_handle, int32_t result_handle);
  int32_t (*unmarshal_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_offset, int32_t length);
  int32_t (*managed_unmarshal_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_handle);
  int32_t (*unmarshal_compressed_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_offset, int32_t length);
  int32_t (*managed_unmarshal_compressed_ec_func_ptr)(void *context, int32_t x_result_handle, int32_t y_result_handle, int32_t ec_handle, int32_t data_handle);
  int32_t (*generate_key_ec_func_ptr)(void *context, int32_t x_pub_key_handle, int32_t y_pub_key_handle, int32_t ec_handle, int32_t result_offset);
  int32_t (*managed_generate_key_ec_func_ptr)(void *context, int32_t x_pub_key_handle, int32_t y_pub_key_handle, int32_t ec_handle, int32_t result_handle);
  int32_t (*create_ec_func_ptr)(void *context, int32_t data_offset, int32_t data_length);
  int32_t (*managed_create_ec_func_ptr)(void *context, int32_t data_handle);
  int32_t (*get_curve_length_ec_func_ptr)(void *context, int32_t ec_handle);
  int32_t (*get_priv_key_byte_length_ec_func_ptr)(void *context, int32_t ec_handle);
  int32_t (*elliptic_curve_get_values_func_ptr)(void *context, int32_t ec_handle, int32_t field_order_handle, int32_t base_point_order_handle, int32_t eq_constant_handle, int32_t x_base_point_handle, int32_t y_base_point_handle);
} vm_exec_vm_hook_pointers;

/**
 * Opaque pointer to a `wasmer_runtime::Ctx` value in Rust.
 *
 * An instance context is passed to any host function (aka imported
 * function) as the first argument. It is necessary to read the
 * instance data or the memory, respectively with the
 * `wasmer_instance_context_data_get()` function, and the
 * `wasmer_instance_context_memory()` function.
 *
 * It is also possible to get the instance context outside a host
 * function by using the `wasmer_instance_context_get()`
 * function. See also `wasmer_instance_context_data_set()` to set the
 * instance context data.
 *
 * Example:
 *
 * ```c
 * // A host function that prints data from the WebAssembly memory to
 * // the standard output.
 * void print(wasmer_instance_context_t *context, int32_t pointer, int32_t length) {
 *     // Use `wasmer_instance_context` to get back the first instance memory.
 *     const wasmer_memory_t *memory = wasmer_instance_context_memory(context, 0);
 *
 *     // Continue…
 * }
 * ```
 */
typedef struct {

} vm_exec_compilation_options_t;

vm_exec_result_t vm_check_signatures(vm_exec_instance_t *_instance);

int vm_exec_execution_info_flush(char *dest_buffer, int dest_buffer_len);

int vm_exec_execution_info_length(void);

void vm_exec_executor_destroy(vm_exec_executor_t *executor);

/**
 * Sets the data that can be hold by an instance context.
 *
 * An instance context (represented by the opaque
 * `wasmer_instance_context_t` structure) can hold user-defined
 * data. This function sets the data. This function is complementary
 * of `wasmer_instance_context_data_get()`.
 *
 * This function does nothing if `instance` is a null pointer.
 */
vm_exec_result_t vm_exec_executor_set_context_ptr(vm_exec_executor_t *executor, void *context_ptr);

/**
 * Calls an exported function of a WebAssembly instance by `name`
 * with the provided parameters. The exported function results are
 * stored on the provided `results` pointer.
 *
 * This function returns `vm_exec_result_t::WASMER_OK` upon success,
 * `vm_exec_result_t::WASMER_ERROR` otherwise. You can use
 * `wasmer_last_error_message()` to get the generated error message.
 *
 * Potential errors are the following:
 *
 *   * `instance` is a null pointer,
 *   * `name` is a null pointer,
 *   * `params` is a null pointer.
 */
vm_exec_result_t vm_exec_instance_call(vm_exec_instance_t *instance, const char *func_name_ptr);

/**
 * Frees memory for the given `vm_exec_instance_t`.
 *
 * Check the `wasmer_instantiate()` function to get a complete
 * example.
 *
 * If `instance` is a null pointer, this function does nothing.
 *
 * Example:
 *
 * ```c
 * // Get an instance.
 * vm_exec_instance_t *instance = NULL;
 * wasmer_instantiate(&instance, bytes, bytes_length, imports, 0);
 *
 * // Destroy the instance.
 * wasmer_instance_destroy(instance);
 * ```
 */
void vm_exec_instance_destroy(vm_exec_instance_t *instance);

int vm_exec_instance_has_function(vm_exec_instance_t *instance, const char *func_name_ptr);

/**
 * Gets the length in bytes of the last error if any.
 *
 * This can be used to dynamically allocate a buffer with the correct number of
 * bytes needed to store a message.
 */
int vm_exec_last_error_length(void);

/**
 * Gets the last error message if any into the provided buffer
 * `buffer` up to the given `length`.
 *
 * The `length` parameter must be large enough to store the last
 * error message. Ideally, the value should come from
 * `wasmer_last_error_length()`.
 *
 * The function returns the length of the string in bytes, `-1` if an
 * error occurs. Potential errors are:
 *
 *  * The buffer is a null pointer,
 *  * The buffer is too smal to hold the error message.
 *
 * Note: The error message always has a trailing null character.
 * ```
 */
int vm_exec_last_error_message(char *dest_buffer, int dest_buffer_len);

vm_exec_result_t vm_exec_new_executor(vm_exec_executor_t **executor,
                                      vm_exec_vm_hook_pointers **vm_hook_pointers_ptr_ptr);

vm_exec_result_t vm_exec_new_instance(vm_exec_executor_t *executor,
                                      vm_exec_instance_t **instance,
                                      uint8_t *wasm_bytes_ptr,
                                      uint32_t wasm_bytes_len,
                                      const vm_exec_compilation_options_t *options_ptr);

int vm_exported_function_names(vm_exec_instance_t *instance,
                               char *dest_buffer,
                               int dest_buffer_len);

int vm_exported_function_names_length(vm_exec_instance_t *instance);
