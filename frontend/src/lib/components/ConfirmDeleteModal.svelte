<script lang="ts">
  import { AlertTriangle, Trash2, Loader2 } from 'lucide-svelte';

  interface Props {
    isOpen: boolean;
    title?: string;
    message?: string;
    itemName?: string;
    confirmText?: string;
    cancelText?: string;
    isDeleting?: boolean;
    onConfirm: () => void | Promise<void>;
    onCancel: () => void;
  }

  let {
    isOpen = false,
    title = 'Konfirmasi Hapus Data',
    message = 'Data yang telah dihapus tidak dapat dipulihkan kembali dari sistem database.',
    itemName = '',
    confirmText = 'Ya, Hapus Data',
    cancelText = 'Batal',
    isDeleting = false,
    onConfirm,
    onCancel
  }: Props = $props();

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === 'Escape' && isOpen && !isDeleting) {
      onCancel();
    }
  }
</script>

<svelte:window onkeydown={handleKeydown} />

{#if isOpen}
  <div
    class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6 bg-slate-950/70 backdrop-blur-sm transition-all"
    role="dialog"
    aria-modal="true"
    aria-labelledby="modal-title"
  >
    <!-- Backdrop dismiss button -->
    <button
      type="button"
      class="fixed inset-0 w-full h-full cursor-default -z-10 bg-transparent border-0 focus:outline-none"
      onclick={() => {
        if (!isDeleting) onCancel();
      }}
      aria-label="Tutup dialog"
      tabindex="-1"
    ></button>

    <div
      class="w-full max-w-md bg-white rounded-3xl shadow-2xl border border-slate-200/90 overflow-hidden transform transition-all animate-in fade-in zoom-in-95 duration-200"
    >
      <div class="p-6 sm:p-8 space-y-5">
        <!-- Icon & Header -->
        <div class="flex items-start gap-4">
          <div class="w-12 h-12 rounded-2xl bg-rose-100/90 text-rose-600 flex items-center justify-center shrink-0 shadow-sm border border-rose-200/50">
            <Trash2 class="w-6 h-6 stroke-[2.2]" />
          </div>
          <div class="space-y-1">
            <h3 id="modal-title" class="font-display font-black text-xl text-slate-950 tracking-tight">
              {title}
            </h3>
            <p class="text-xs text-rose-600 font-semibold uppercase tracking-wider flex items-center gap-1.5">
              <AlertTriangle class="w-3.5 h-3.5" />
              Tindakan Permanen & Destruktif
            </p>
          </div>
        </div>

        <!-- Details -->
        <div class="space-y-3 bg-slate-50 p-4 rounded-2xl border border-slate-100 text-sm">
          {#if itemName}
            <div class="font-semibold text-slate-900 break-words line-clamp-2">
              "{itemName}"
            </div>
          {/if}
          <p class="text-xs sm:text-sm text-slate-600 leading-relaxed font-normal">
            {message}
          </p>
        </div>

        <!-- Action Buttons -->
        <div class="pt-2 flex flex-col-reverse sm:flex-row items-center gap-3">
          <button
            type="button"
            disabled={isDeleting}
            onclick={onCancel}
            class="w-full sm:flex-1 py-3 px-4 rounded-xl border border-slate-300 hover:bg-slate-100 text-slate-700 text-sm font-semibold transition active:scale-95 disabled:opacity-50"
          >
            {cancelText}
          </button>

          <button
            type="button"
            disabled={isDeleting}
            onclick={onConfirm}
            class="w-full sm:flex-1 py-3 px-4 rounded-xl bg-gradient-to-r from-rose-600 to-red-600 hover:from-rose-500 hover:to-red-500 text-white text-sm font-bold shadow-lg shadow-rose-600/25 transition active:scale-95 flex items-center justify-center gap-2 disabled:opacity-60"
          >
            {#if isDeleting}
              <Loader2 class="w-4 h-4 animate-spin" />
              <span>Menghapus...</span>
            {:else}
              <Trash2 class="w-4 h-4" />
              <span>{confirmText}</span>
            {/if}
          </button>
        </div>
      </div>
    </div>
  </div>
{/if}

