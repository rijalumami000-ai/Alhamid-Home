<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { ContactMessage } from '$lib/types';
  import { Mail, Check, Trash2, X, Phone, MessageSquare } from 'lucide-svelte';

  let list = $state<ContactMessage[]>([]);
  let selectedMessage = $state<ContactMessage | null>(null);

  async function loadData() {
    try {
      list = await api.admin.getMessages();
    } catch (e) {
      console.error(e);
    }
  }

  async function markRead(msg: ContactMessage) {
    try {
      await api.admin.markMessageRead(msg.id);
      msg.is_read = true;
      await loadData();
    } catch (e) {
      console.error(e);
    }
  }

  async function handleDelete(id: number) {
    if (confirm('Hapus pesan ini dari kontak masuk?')) {
      try {
        await api.admin.deleteMessage(id);
        if (selectedMessage?.id === id) selectedMessage = null;
        await loadData();
      } catch (e) {
        console.error(e);
      }
    }
  }

  function viewMessage(msg: ContactMessage) {
    selectedMessage = msg;
    if (!msg.is_read) {
      markRead(msg);
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Pesan Masuk Pengunjung | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <div>
    <h1 class="font-serif-luxury font-bold text-2xl text-slate-900">
      Kotak Masuk Pesan & Pertanyaan
    </h1>
    <p class="text-xs text-slate-500 mt-1">
      Pesan yang dikirimkan oleh calon wali santri dan masyarakat melalui formulir kontak.
    </p>
  </div>

  <div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
    <!-- Messages List (5 cols) -->
    <div class="lg:col-span-5 bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-sm flex flex-col h-[75vh]">
      <div class="p-4 border-b border-slate-100 bg-slate-50 text-xs font-bold text-slate-600 uppercase">
        Semua Pesan ({list.length})
      </div>
      <div class="flex-1 overflow-y-auto divide-y divide-slate-100">
        {#each list as msg}
          <div
            class="p-4 cursor-pointer hover:bg-slate-50 transition space-y-1 {selectedMessage?.id === msg.id ? 'bg-emerald-50/70 border-l-4 border-emerald-700' : ''}"
            role="button"
            tabindex="0"
            onclick={() => viewMessage(msg)}
            onkeydown={(e) => e.key === 'Enter' && viewMessage(msg)}
          >
            <div class="flex items-center justify-between">
              <span class="font-bold text-xs text-slate-900 {msg.is_read ? 'font-normal text-slate-600' : 'font-bold text-slate-900'}">
                {msg.name}
              </span>
              <span class="text-[10px] text-slate-400">
                {new Date(msg.created_at).toLocaleDateString('id-ID')}
              </span>
            </div>
            <div class="text-xs font-semibold text-slate-800 line-clamp-1">{msg.subject}</div>
            <p class="text-[11px] text-slate-500 line-clamp-1">{msg.message}</p>
          </div>
        {/each}
        {#if list.length === 0}
          <div class="p-8 text-center text-xs text-slate-400">
            Belum ada pesan yang masuk.
          </div>
        {/if}
      </div>
    </div>

    <!-- Message Detail (7 cols) -->
    <div class="lg:col-span-7 bg-white rounded-3xl border border-slate-200 p-8 shadow-sm flex flex-col justify-between h-[75vh]">
      {#if selectedMessage}
        <div class="space-y-6 overflow-y-auto">
          <div class="flex items-start justify-between pb-4 border-b border-slate-100">
            <div>
              <span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold uppercase {selectedMessage.is_read ? 'bg-emerald-100 text-emerald-800' : 'bg-amber-100 text-amber-800'}">
                {selectedMessage.is_read ? 'Sudah Dibaca' : 'Baru'}
              </span>
              <h2 class="font-serif-luxury font-bold text-xl text-slate-900 mt-2">
                {selectedMessage.subject}
              </h2>
              <div class="text-xs text-slate-500 mt-1">
                Dari: <span class="font-bold text-slate-700">{selectedMessage.name}</span> ({selectedMessage.email})
              </div>
              {#if selectedMessage.phone}
                <div class="text-xs text-slate-500 flex items-center gap-1 mt-0.5">
                  <Phone class="w-3.5 h-3.5 text-emerald-600" />
                  <span>{selectedMessage.phone}</span>
                </div>
              {/if}
            </div>

            <button
              onclick={() => handleDelete(selectedMessage.id)}
              class="p-2 rounded-xl text-slate-400 hover:text-red-600 hover:bg-red-50"
              title="Hapus Pesan"
            >
              <Trash2 class="w-5 h-5" />
            </button>
          </div>

          <div class="space-y-2">
            <div class="text-xs font-bold uppercase text-slate-400">Isi Pesan:</div>
            <p class="text-slate-800 text-sm leading-relaxed whitespace-pre-line p-4 rounded-2xl bg-slate-50 border border-slate-100">
              {selectedMessage.message}
            </p>
          </div>
        </div>

        <!-- Quick Reply Options -->
        <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
          <div class="text-xs text-slate-400">
            Diterima: {new Date(selectedMessage.created_at).toLocaleString('id-ID')}
          </div>
          <div class="flex items-center gap-2">
            {#if selectedMessage.phone}
              <a
                href="https://wa.me/{selectedMessage.phone.replace(/[^0-9]/g, '')}"
                target="_blank"
                class="px-4 py-2 rounded-xl bg-emerald-800 text-white text-xs font-bold hover:bg-emerald-700 transition"
              >
                Balas via WhatsApp
              </a>
            {/if}
            <a
              href="mailto:{selectedMessage.email}?subject=Balasan: {encodeURIComponent(selectedMessage.subject)}"
              class="px-4 py-2 rounded-xl border border-slate-200 text-slate-700 text-xs font-bold hover:bg-slate-50 transition"
            >
              Balas via Email
            </a>
          </div>
        </div>
      {:else}
        <div class="h-full flex flex-col items-center justify-center text-slate-400 space-y-2">
          <MessageSquare class="w-10 h-10 text-slate-300" />
          <p class="text-xs">Pilih salah satu pesan di sebelah kiri untuk membaca isinya.</p>
        </div>
      {/if}
    </div>
  </div>
</div>

