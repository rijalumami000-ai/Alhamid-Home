<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Article, Category } from '$lib/types';
  import { Plus, Edit2, Trash2, X, Check, Search, Eye, Sparkles } from 'lucide-svelte';
  import ConfirmDeleteModal from '$lib/components/ConfirmDeleteModal.svelte';
  import DriveImageInput from '$lib/components/DriveImageInput.svelte';

  let activeTab = $state<'articles' | 'categories'>('articles');
  let articles = $state<Article[]>([]);
  let categories = $state<Category[]>([]);
  let loading = $state(true);

  // Article Modal
  let isArticleModalOpen = $state(false);
  let editingArticle = $state<Article | null>(null);
  let articleForm = $state({
    title: '',
    category_id: 1,
    excerpt: '',
    content: '',
    thumbnail_url: '',
    author: 'Redaksi Al-Hamid',
    status: 'published' as 'published' | 'draft',
    is_featured: false,
  });

  // Category Modal
  let isCategoryModalOpen = $state(false);
  let editingCategory = $state<Category | null>(null);
  let categoryForm = $state({
    name: '',
    slug: '',
    type: 'berita',
    description: '',
  });

  // Delete modal state
  let isDeleteModalOpen = $state(false);
  let deleteTarget = $state<{ type: 'article' | 'category'; id: number; name: string } | null>(null);
  let isDeleting = $state(false);

  async function loadData() {
    loading = true;
    try {
      const [art, cat] = await Promise.all([
        api.admin.getArticles(),
        api.admin.getCategories(),
      ]);
      articles = art;
      categories = cat;
      if (cat.length > 0 && !editingArticle) {
        articleForm.category_id = cat[0].id;
      }
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  function openCreateArticle() {
    editingArticle = null;
    articleForm = {
      title: '',
      category_id: categories[0]?.id || 1,
      excerpt: '',
      content: '',
      thumbnail_url: 'https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80',
      author: 'Redaksi Al-Hamid',
      status: 'published',
      is_featured: false,
    };
    isArticleModalOpen = true;
  }

  function openEditArticle(art: Article) {
    editingArticle = art;
    articleForm = {
      title: art.title,
      category_id: art.category_id,
      excerpt: art.excerpt,
      content: art.content,
      thumbnail_url: art.thumbnail_url,
      author: art.author,
      status: art.status,
      is_featured: art.is_featured,
    };
    isArticleModalOpen = true;
  }

  async function handleSaveArticle(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editingArticle) {
        await api.admin.updateArticle(editingArticle.id, articleForm);
      } else {
        await api.admin.createArticle(articleForm);
      }
      isArticleModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan artikel');
    }
  }

  function promptDeleteArticle(art: Article) {
    deleteTarget = { type: 'article', id: art.id, name: art.title };
    isDeleteModalOpen = true;
  }

  // Category Actions
  function openCreateCategory() {
    editingCategory = null;
    categoryForm = { name: '', slug: '', type: 'berita', description: '' };
    isCategoryModalOpen = true;
  }

  function openEditCategory(cat: Category) {
    editingCategory = cat;
    categoryForm = { name: cat.name, slug: cat.slug, type: cat.type, description: cat.description || '' };
    isCategoryModalOpen = true;
  }

  async function handleSaveCategory(e: SubmitEvent) {
    e.preventDefault();
    try {
      if (editingCategory) {
        await api.admin.updateCategory(editingCategory.id, categoryForm as any);
      } else {
        await api.admin.createCategory(categoryForm as any);
      }
      isCategoryModalOpen = false;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menyimpan kategori');
    }
  }

  function promptDeleteCategory(cat: Category) {
    deleteTarget = { type: 'category', id: cat.id, name: cat.name };
    isDeleteModalOpen = true;
  }

  async function handleConfirmDelete() {
    if (!deleteTarget) return;
    isDeleting = true;
    try {
      if (deleteTarget.type === 'article') {
        await api.admin.deleteArticle(deleteTarget.id);
      } else {
        await api.admin.deleteCategory(deleteTarget.id);
      }
      isDeleteModalOpen = false;
      deleteTarget = null;
      await loadData();
    } catch (err: any) {
      alert(err.message || 'Gagal menghapus data');
    } finally {
      isDeleting = false;
    }
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Kelola Artikel & Berita | CMS Al-Hamid</title>
</svelte:head>

<div class="space-y-6">
  <!-- Header -->
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="font-display font-black text-2xl text-slate-900 tracking-tight">
        Artikel, Berita & Kategori
      </h1>
      <p class="text-xs text-slate-500 mt-1">
        Publikasikan kajian kitab kuning, informasi pesantren, dan pengumuman resmi.
      </p>
    </div>
    <div class="flex items-center gap-2">
      {#if activeTab === 'articles'}
        <button
          onclick={openCreateArticle}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
        >
          <Plus class="w-4 h-4" /> Tulis Artikel Baru
        </button>
      {:else}
        <button
          onclick={openCreateCategory}
          class="inline-flex items-center gap-1.5 px-4 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition active:scale-95"
        >
          <Plus class="w-4 h-4" /> Tambah Kategori
        </button>
      {/if}
    </div>
  </div>

  <!-- Tabs Navigation -->
  <div class="flex items-center gap-2 border-b border-slate-200">
    <button
      onclick={() => (activeTab = 'articles')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'articles'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Daftar Artikel ({articles.length})
    </button>
    <button
      onclick={() => (activeTab = 'categories')}
      class="px-5 py-3 text-xs font-bold border-b-2 transition {activeTab === 'categories'
        ? 'border-teal-700 text-teal-800'
        : 'border-transparent text-slate-500 hover:text-slate-800'}"
    >
      Daftar Kategori ({categories.length})
    </button>
  </div>

  <!-- ARTICLES LIST -->
  {#if activeTab === 'articles'}
    <div class="bg-white rounded-3xl border border-slate-200 shadow-sm overflow-hidden">
      {#if articles.length === 0}
        <div class="p-12 text-center text-slate-400 text-sm">Belum ada artikel yang ditambahkan.</div>
      {:else}
        <div class="overflow-x-auto">
          <table class="w-full text-left text-xs text-slate-600">
            <thead class="bg-slate-50 border-b border-slate-200 text-slate-700 font-bold uppercase tracking-wider text-[10px]">
              <tr>
                <th class="py-3.5 px-4">Artikel</th>
                <th class="py-3.5 px-4">Kategori</th>
                <th class="py-3.5 px-4">Penulis</th>
                <th class="py-3.5 px-4">Status</th>
                <th class="py-3.5 px-4">Tanggal</th>
                <th class="py-3.5 px-4 text-right">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100">
              {#each articles as art}
                <tr class="hover:bg-slate-50/60 transition">
                  <td class="py-3.5 px-4 flex items-center gap-3">
                    <div class="w-12 h-12 rounded-xl bg-slate-100 overflow-hidden shrink-0 border border-slate-200">
                      {#if art.thumbnail_url}
                        <img src={art.thumbnail_url} alt={art.title} class="w-full h-full object-cover" />
                      {:else}
                        <div class="w-full h-full flex items-center justify-center text-slate-400">Foto</div>
                      {/if}
                    </div>
                    <div class="space-y-0.5 max-w-sm">
                      <div class="font-bold text-slate-900 line-clamp-1">{art.title}</div>
                      <div class="text-[11px] text-slate-400 font-mono">/{art.slug}</div>
                      {#if art.is_featured}
                        <span class="inline-flex items-center gap-1 text-[9px] font-bold text-amber-700 bg-amber-50 px-1.5 py-0.2 rounded border border-amber-200">
                          <Sparkles class="w-2.5 h-2.5" /> Featured
                        </span>
                      {/if}
                    </div>
                  </td>
                  <td class="py-3.5 px-4 font-medium text-slate-800">
                    {art.category?.name || '-'}
                  </td>
                  <td class="py-3.5 px-4 text-slate-600 font-medium">{art.author}</td>
                  <td class="py-3.5 px-4">
                    <span
                      class="px-2 py-0.5 rounded-full text-[10px] font-bold {art.status === 'published'
                        ? 'bg-emerald-100 text-emerald-800'
                        : 'bg-slate-100 text-slate-600'}"
                    >
                      {art.status === 'published' ? 'Tayang' : 'Draf'}
                    </span>
                  </td>
                  <td class="py-3.5 px-4 font-mono text-[11px] text-slate-500">
                    {art.published_at ? new Date(art.published_at).toLocaleDateString('id-ID') : '-'}
                  </td>
                  <td class="py-3.5 px-4 text-right">
                    <div class="inline-flex items-center gap-1">
                      <a
                        href="/berita/{art.slug}"
                        target="_blank"
                        class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
                        title="Lihat Pratinjau"
                      >
                        <Eye class="w-4 h-4" />
                      </a>
                      <button
                        onclick={() => openEditArticle(art)}
                        class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
                        title="Edit Artikel"
                      >
                        <Edit2 class="w-4 h-4" />
                      </button>
                      <button
                        onclick={() => promptDeleteArticle(art)}
                        class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
                        title="Hapus Artikel"
                      >
                        <Trash2 class="w-4 h-4" />
                      </button>
                    </div>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      {/if}
    </div>
  {/if}

  <!-- CATEGORIES LIST -->
  {#if activeTab === 'categories'}
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
      {#each categories as cat}
        <div class="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-3 flex flex-col justify-between">
          <div class="space-y-1">
            <div class="flex items-center justify-between">
              <h3 class="font-bold text-slate-900 text-sm">{cat.name}</h3>
              <span class="text-[10px] uppercase font-mono px-2 py-0.5 bg-slate-100 rounded text-slate-600 font-semibold">
                {cat.type}
              </span>
            </div>
            <p class="text-xs text-slate-400 font-mono">slug: {cat.slug}</p>
            <p class="text-xs text-slate-600 line-clamp-2">{cat.description || 'Tidak ada deskripsi'}</p>
          </div>

          <div class="pt-3 border-t border-slate-100 flex items-center justify-end gap-2">
            <button
              onclick={() => openEditCategory(cat)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-teal-700 hover:bg-slate-100 transition"
              title="Edit Kategori"
            >
              <Edit2 class="w-4 h-4" />
            </button>
            <button
              onclick={() => promptDeleteCategory(cat)}
              class="p-1.5 rounded-lg text-slate-500 hover:text-rose-600 hover:bg-rose-50 transition"
              title="Hapus Kategori"
            >
              <Trash2 class="w-4 h-4" />
            </button>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<!-- ARTICLE FORM MODAL -->
{#if isArticleModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-2xl rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editingArticle ? 'Edit Artikel' : 'Tulis Artikel Baru'}
        </h2>
        <button onclick={() => (isArticleModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSaveArticle} class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
        <div class="space-y-1">
          <label for="art_title" class="text-xs font-bold text-slate-700 uppercase">Judul Artikel *</label>
          <input
            id="art_title"
            type="text"
            required
            bind:value={articleForm.title}
            placeholder="Judul artikel atau berita..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-teal-600/30"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1">
            <label for="art_cat" class="text-xs font-bold text-slate-700 uppercase">Kategori *</label>
            <select
              id="art_cat"
              bind:value={articleForm.category_id}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-teal-600/30 font-medium bg-white"
            >
              {#each categories as c}
                <option value={c.id}>{c.name}</option>
              {/each}
            </select>
          </div>
          <div class="space-y-1">
            <label for="art_author" class="text-xs font-bold text-slate-700 uppercase">Penulis / Sumber</label>
            <input
              id="art_author"
              type="text"
              bind:value={articleForm.author}
              placeholder="Redaksi Al-Hamid"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-teal-600/30 font-medium"
            />
          </div>
        </div>

        <DriveImageInput
          id="art_thumb"
          label="Thumbnail Foto Artikel"
          bind:value={articleForm.thumbnail_url}
          placeholder="https://drive.google.com/... atau URL gambar"
        />

        <div class="space-y-1">
          <label for="art_excerpt" class="text-xs font-bold text-slate-700 uppercase">Ringkasan Singkat (Excerpt)</label>
          <textarea
            id="art_excerpt"
            rows="2"
            bind:value={articleForm.excerpt}
            placeholder="Ringkasan 1-2 kalimat untuk kartu depan..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-teal-600/30 font-medium"
          ></textarea>
        </div>

        <div class="space-y-1">
          <label for="art_content" class="text-xs font-bold text-slate-700 uppercase">Isi Lengkap Artikel (Mendukung Format HTML) *</label>
          <textarea
            id="art_content"
            required
            rows="8"
            bind:value={articleForm.content}
            placeholder="Tuliskan naskah artikel, paragraf, <h3>, <ul>, dsb..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm focus:outline-none focus:ring-2 focus:ring-teal-600/30 font-mono text-xs leading-relaxed"
          ></textarea>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-2">
          <div class="space-y-1">
            <label for="art_status" class="text-xs font-bold text-slate-700 uppercase">Status Publikasi</label>
            <select
              id="art_status"
              bind:value={articleForm.status}
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm bg-white font-medium"
            >
              <option value="published">Tayang (Published)</option>
              <option value="draft">Draf (Draft)</option>
            </select>
          </div>
          <div class="flex items-center gap-2 pt-6">
            <input
              id="art_featured"
              type="checkbox"
              bind:checked={articleForm.is_featured}
              class="w-4 h-4 text-teal-600 rounded"
            />
            <label for="art_featured" class="text-xs font-bold text-slate-700">Tampilkan sebagai Berita Unggulan (Featured)</label>
          </div>
        </div>

        <div class="pt-6 border-t border-slate-100 flex items-center justify-end gap-3">
          <button
            type="button"
            onclick={() => (isArticleModalOpen = false)}
            class="px-5 py-2.5 rounded-xl border border-slate-200 text-slate-600 text-xs font-bold hover:bg-slate-50"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-6 py-2.5 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Artikel
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CATEGORY FORM MODAL -->
{#if isCategoryModalOpen}
  <div class="fixed inset-0 z-50 bg-slate-950/60 backdrop-blur-sm flex items-center justify-center p-4">
    <div class="bg-white w-full max-w-md rounded-3xl shadow-2xl border border-slate-200 overflow-hidden">
      <div class="p-6 border-b border-slate-100 flex items-center justify-between">
        <h2 class="font-display font-bold text-lg text-slate-900">
          {editingCategory ? 'Edit Kategori' : 'Tambah Kategori'}
        </h2>
        <button onclick={() => (isCategoryModalOpen = false)}>
          <X class="w-5 h-5 text-slate-400" />
        </button>
      </div>

      <form onsubmit={handleSaveCategory} class="p-6 space-y-4">
        <div class="space-y-1">
          <label for="cat_name" class="text-xs font-bold text-slate-700 uppercase">Nama Kategori *</label>
          <input
            id="cat_name"
            type="text"
            required
            bind:value={categoryForm.name}
            placeholder="Contoh: Kajian Fiqih"
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          />
        </div>
        <div class="space-y-1">
          <label for="cat_desc" class="text-xs font-bold text-slate-700 uppercase">Deskripsi</label>
          <textarea
            id="cat_desc"
            rows="2"
            bind:value={categoryForm.description}
            placeholder="Keterangan kategori..."
            class="w-full px-4 py-2.5 rounded-xl border border-slate-200 text-sm font-medium"
          ></textarea>
        </div>
        <div class="pt-4 border-t border-slate-100 flex justify-end gap-2">
          <button
            type="button"
            onclick={() => (isCategoryModalOpen = false)}
            class="px-4 py-2 rounded-xl border border-slate-200 text-xs font-bold text-slate-600"
          >
            Batal
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-teal-700 hover:bg-teal-800 text-white text-xs font-bold shadow-md transition"
          >
            Simpan Kategori
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- CONFIRM DELETE MODAL -->
<ConfirmDeleteModal
  isOpen={isDeleteModalOpen}
  title={deleteTarget?.type === 'article' ? 'Hapus Artikel Berita' : 'Hapus Kategori'}
  itemName={deleteTarget?.name}
  message={deleteTarget?.type === 'article'
    ? 'Artikel ini akan dihapus secara permanen dari sistem CMS dan publikasi web.'
    : 'Kategori ini akan dihapus. Periksa kembali jika ada artikel yang mengacu pada kategori ini.'}
  {isDeleting}
  onConfirm={handleConfirmDelete}
  onCancel={() => {
    isDeleteModalOpen = false;
    deleteTarget = null;
  }}
/>
